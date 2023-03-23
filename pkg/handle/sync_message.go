package handle

import (
	"log"
	"os"
	"time"

	"fsm_client/pkg/ent"
	"fsm_client/pkg/ignore"
)

func (h *Handle) FileChange(action string, file ent.File, parentPath, rootPath string) {

	if h.Ignore.Match(file.Name) {
		return
	}

	log.Println(file.Name, action)
	ignore.Lock.Store(rootPath+parentPath+file.Name, 1)

	switch action {
	case "update", "create":
		err := h.FileWrite(file, parentPath, rootPath)
		if err != nil {
			log.Println(err)
		}
	case "delete":
		err := h.FileDelete(file, parentPath, rootPath)
		if err != nil {
			log.Println(err)
		}
	case "rename":
		log.Println("文件更名")

		var f ent.File
		h.DB.Where("sync_id = ? and id = ? and level =?", file.SyncID, file.ID, file.Level).Find(&f)
		ignore.Lock.Store(rootPath+parentPath+f.Name, 1)

		if err := os.Rename(rootPath+parentPath+f.Name, rootPath+parentPath+file.Name); err != nil {
			log.Println(err)
		}

		h.DB.Save(&file)
		time.Sleep(time.Millisecond * 500)
		ignore.Lock.Delete(rootPath + parentPath + f.Name)
	}

	if action != "rename" {
		time.Sleep(time.Millisecond * 500)
	}
	ignore.Lock.Delete(rootPath + parentPath + file.Name)
}
