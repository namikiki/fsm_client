package handle

import (
	"io"
	"os"
	"path/filepath"

	"fsm_client/pkg/ent"
)

func (h *Handle) FIleCreate(file ent.File, parentPath, rootPath string) error {

	fileIO, err := h.HttpClient.GetFile(file.ID)
	if err != nil {
		return err
	}

	createFile, err := os.Create(filepath.Join(rootPath, parentPath, file.Name))
	if err != nil {
		return err
	}

	defer fileIO.Close()
	defer createFile.Close()
	_, err = io.Copy(createFile, fileIO)
	h.DB.Create(&file)
	return err
}

func (h *Handle) FileDelete(file ent.File, parentPath, rootPath string) error {
	err := os.Remove(filepath.Join(rootPath, parentPath, file.Name))
	h.DB.Delete(&file)
	return err
}
