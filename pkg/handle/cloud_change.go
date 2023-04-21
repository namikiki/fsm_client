package handle

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"fsm_client/pkg/ent"
	"fsm_client/pkg/ignore"
)

func (h *Handle) FileChange(action string, file ent.File, parentPath, rootPath string) {

	if h.Ignore.Match(file.Name) {
		return
	}

	filePath := filepath.Join(rootPath, parentPath, file.Name)
	log.Println(filePath, action)
	ignore.Lock.Store(filePath, 1)

	switch action {
	case "update", "create":
		if err := h.FileWrite(file, filePath); err != nil {
			log.Println(err)
		}
	case "delete":
		if err := h.FileDelete(file, filePath); err != nil {
			log.Println(err)
		}
	case "rename":
		log.Println("文件更名")

		var f ent.File
		h.DB.Where("sync_id = ? and id = ? and level =?", file.SyncID, file.ID, file.Level).Find(&f)
		oldFilePath := filepath.Join(rootPath, parentPath, f.Name)
		ignore.Lock.Store(oldFilePath, 1)

		if err := os.Rename(oldFilePath, filePath); err != nil {
			log.Println(err)
		}

		h.DB.Save(&file)
		time.Sleep(time.Millisecond * 500)
		ignore.Lock.Delete(oldFilePath)
	}

	if action != "rename" {
		time.Sleep(time.Millisecond * 500)
	}
	ignore.Lock.Delete(filePath)
}

func (h *Handle) DirChange(action string, dir ent.Dir, rootPath string) {

	dirPath := filepath.Join(rootPath, dir.Dir)
	ignore.Lock.Store(dirPath, 1)

	switch action {
	case "create":

		if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
			return
		}
		h.DB.Create(&dir)

	case "delete":

		if err := os.RemoveAll(dirPath); err != nil {
			return
		}
		h.DB.Delete(&dir)

	case "rename":

		var d ent.Dir
		h.DB.Where("sync_id =? and id =? and level = ?", dir.SyncID, dir.ID, dir.Level).Find(&d)

		oldDirPath := filepath.Join(rootPath, d.Dir)
		ignore.Lock.Store(oldDirPath, 1)

		if err := os.Rename(oldDirPath, dirPath); err != nil {
			return
		}
		h.DB.Save(&dir)

		time.Sleep(time.Millisecond * 500)
		ignore.Lock.Delete(oldDirPath)
	}

	if action != "rename" {
		time.Sleep(time.Millisecond * 500)
	}

	ignore.Lock.Delete(dirPath)
}

func (h *Handle) SyncTaskCreate(synctask ent.SyncTask) {
	h.DB.Create(&synctask)
}

func (h *Handle) SyncTaskDelete(synctask ent.SyncTask) {

	h.DB.Delete(&synctask)
}
