package handle

import (
	"log"

	"fsm_client/pkg/ent"
)

func (h *Handle) FileChange(action string, file ent.File, parentPath, rootPath string) {

	if h.Ignore.Match(file.Name) {
		return
	}

	log.Println(file.Name, action)

	if action == "update" || action == "create" {

		err := h.FileWrite(file, parentPath, rootPath)
		if err != nil {
			log.Println(err)
		}
	} else {
		err := h.FileDelete(file, parentPath, rootPath)
		if err != nil {
			log.Println(err)
		}
	}
}
