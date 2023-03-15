package handle

import (
	"os"
	"path/filepath"

	"fsm_client/pkg/ent"
)

func (h *Handle) DirCreate(dir ent.Dir, rootPath string) {
	os.MkdirAll(filepath.Join(rootPath, dir.Dir), os.ModePerm)
	h.DB.Create(&dir)
}

func (h *Handle) DirDelete(dir ent.Dir, rootPath string) {
	os.RemoveAll(filepath.Join(rootPath, dir.Dir))
	h.DB.Delete(&dir)
}
