package handle

import (
	"io"
	"os"
	"time"

	"fsm_client/pkg/ent"
)

const PathSeparator = "/"

func (h *Handle) FileWrite(file ent.File, filepath string) error {
	fileIO, err := h.HttpClient.GetFile(file.ID)
	if err != nil {
		return err
	}

	//TODO 文件修改冲突
	createFile, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	_, err = io.Copy(createFile, fileIO)

	fileIO.Close()
	createFile.Close()

	if err := os.Chtimes(filepath, time.Unix(file.ModTime, 0), time.Unix(file.ModTime, 0)); err != nil {
		return err
	}

	h.DB.Save(&file)
	return err
}

func (h *Handle) FileDelete(file ent.File, filepath string) error {
	err := os.Remove(filepath)
	h.DB.Delete(&file)
	return err
}

func (h *Handle) DeleteAllFileByDir(path string) error {
	return os.RemoveAll(path)
}
