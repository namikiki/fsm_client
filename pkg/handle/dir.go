package handle

import (
	"os"
	"path/filepath"
	"strings"
	"time"

	"fsm_client/pkg/ent"
	"fsm_client/pkg/ignore"
)

func (h *Handle) DirCreate(dir ent.Dir, rootPath string) {
	ignore.Lock.Store(strings.TrimSuffix(filepath.Join(rootPath, dir.Dir), PathSeparator), 1)

	os.MkdirAll(filepath.Join(rootPath, dir.Dir), os.ModePerm)
	h.DB.Create(&dir)

	time.Sleep(time.Millisecond * 500)
	ignore.Lock.Delete(strings.TrimSuffix(rootPath+dir.Dir, PathSeparator))
}

func (h *Handle) DirDelete(dir ent.Dir, rootPath string) {
	ignore.Lock.Store(strings.TrimSuffix(filepath.Join(rootPath, dir.Dir), PathSeparator), 1)

	os.RemoveAll(filepath.Join(rootPath, dir.Dir))
	h.DB.Delete(&dir)

	time.Sleep(time.Millisecond * 500)
	ignore.Lock.Delete(strings.TrimSuffix(rootPath+dir.Dir, PathSeparator))
}
