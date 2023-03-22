package handle

import (
	"os"
	"path/filepath"
	"strings"
	"time"

	"fsm_client/pkg/ent"
	"fsm_client/pkg/ignore"
)

func (h *Handle) DirChange(action string, dir ent.Dir, rootPath string) {

	ignore.Lock.Store(strings.TrimSuffix(filepath.Join(rootPath, dir.Dir), PathSeparator), 1)

	switch action {
	case "create":

		if err := os.MkdirAll(filepath.Join(rootPath, dir.Dir), os.ModePerm); err != nil {
			return
		}
		h.DB.Create(&dir)

	case "delete":

		if err := os.RemoveAll(filepath.Join(rootPath, dir.Dir)); err != nil {
			return
		}
		h.DB.Delete(&dir)

	case "rename":

		var d ent.Dir
		h.DB.Where("sync_id =? and id =? and level = ?", dir.SyncID, dir.ID, dir.Level).Find(&d)
		ignore.Lock.Store(strings.TrimSuffix(filepath.Join(rootPath, d.Dir), PathSeparator), 1)

		if err := os.Rename(filepath.Join(rootPath, d.Dir),
			filepath.Join(rootPath, dir.Dir)); err != nil {
			return
		}
		h.DB.Save(&dir)

		time.Sleep(time.Millisecond * 500)
		ignore.Lock.Delete(strings.TrimSuffix(filepath.Join(rootPath, d.Dir), PathSeparator))
		ignore.Lock.Delete(strings.TrimSuffix(rootPath+dir.Dir, PathSeparator))
	}

	if action != "rename" {
		time.Sleep(time.Millisecond * 500)
		ignore.Lock.Delete(strings.TrimSuffix(rootPath+dir.Dir, PathSeparator))
	}
}
