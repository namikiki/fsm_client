package handle

import (
	"encoding/json"
	"os"
	"path/filepath"

	"fsm_client/pkg/ent"
)

func (h *Handle) DirCreate(data []byte, rootPath string) {
	var dir ent.Dir
	json.Unmarshal(data, &dir)
	os.MkdirAll(filepath.Join(rootPath, dir.Dir), os.ModePerm)
	h.DB.Create(&dir)
}

func (h *Handle) DirDelete(data []byte, rootPath string) {
	var dir ent.Dir
	json.Unmarshal(data, &dir)
	os.RemoveAll(filepath.Join(rootPath, dir.Dir))
	h.DB.Delete(&dir)
}
