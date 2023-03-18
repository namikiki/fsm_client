package handle

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"fsm_client/pkg/ent"
	fsn "fsm_client/pkg/fsnotify"

	"github.com/fsnotify/fsnotify"
)

const PathSeparator = string(os.PathSeparator)

func (h *Handle) FileWrite(file ent.File, parentPath, rootPath string) error {

	fileIO, err := h.HttpClient.GetFile(file.ID)
	if err != nil {
		return err
	}

	createFile, err := os.OpenFile(filepath.Join(rootPath, parentPath, file.Name), os.O_CREATE|os.O_WRONLY, 0644)
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

func (h *Handle) DeleteAllFileByDir(path string) error {
	return os.RemoveAll(path)
}

func (h *Handle) GetFileParentID(absPath string, fileName string, level int) string {
	suffix := strings.TrimSuffix(absPath, fileName)
	var dir ent.Dir
	h.DB.Where("dir = ? and level = ?", suffix, level-1).Find(&dir)
	return dir.ID
}

func (h *Handle) CloudFileCreate(fw fsn.FsEventWithID, stat os.FileInfo) error {
	level := len(strings.Split(fw.AbsPath, PathSeparator))

	file := ent.File{
		SyncID:     fw.SyncID,
		Name:       stat.Name(),
		Level:      uint64(level),
		Deleted:    false,
		CreateTime: time.Now(),
		ModTime:    stat.ModTime(),
	}

	if file.ParentDirID = h.GetFileParentID(fw.AbsPath, stat.Name(), level); file.ParentDirID == "" {
		return errors.New("未找到 parentID")
	}

	fileIO, err := os.Open(fw.Name)
	if err != nil {
		return err
	}

	if err := h.HttpClient.FileCreate(&file, fileIO); err != nil {
		return err
	}

	h.DB.Create(&file)
	return nil
}

func (h *Handle) CloudFileDelete(file ent.File) error {
	err := h.HttpClient.FileDelete(file)
	if err != nil {
		return err
	}

	h.DB.Delete(&file)
	return nil
}

func (h *Handle) CloudDirCreate(fw fsn.FsEventWithID, stat os.FileInfo) error {
	level := len(strings.Split(fw.AbsPath, PathSeparator))

	dir := ent.Dir{
		SyncID:     fw.SyncID,
		Dir:        fw.AbsPath,
		Level:      uint64(level),
		Deleted:    false,
		CreateTime: time.Now(),
		ModTime:    stat.ModTime(),
	}

	if err := h.HttpClient.DirCreate(&dir); err != nil {
		return err
	}

	h.DB.Create(&dir)
	return nil
}

func (h *Handle) CloudDirDelete(dir ent.Dir) error {
	err := h.HttpClient.DirDelete(dir)
	if err != nil {
		return err
	}

	h.DB.Delete(&dir)
	return nil
}

func (h *Handle) GetDeleteID(fw fsn.FsEventWithID) (*ent.Dir, *ent.File, bool) {
	split := strings.Split(fw.AbsPath, PathSeparator)
	level := len(split)
	var dir ent.Dir
	h.DB.Where("sync_id = ? and level =? dir =?", fw.SyncID, level, fw.AbsPath).Find(&dir)
	if dir.ID != "" {
		return &dir, nil, true
	}

	var files []ent.File
	h.DB.Where("sync_id = ? and level =? name =?", fw.SyncID, level, split[level-1]).Find(&files)
	if len(files) == 1 {
		return nil, &files[0], false
	}

	return nil, nil, false
	//for _, file := range files {
	//
	//	os.Stat()
	//}
}

func (h *Handle) PressLocalChange(eventChan chan fsn.FsEventWithID, errChan chan error) {
	for {
		e := <-eventChan
		switch e.Op {
		case fsnotify.Create:
			stat, err := os.Stat(e.Name)
			if err != nil {
				errChan <- err
				continue
			}

			if stat.IsDir() {
				errChan <- h.CloudDirCreate(e, stat)
				continue
			}
			errChan <- h.CloudFileCreate(e, stat)
		case fsnotify.Write:
			stat, err := os.Stat(e.Name)
			if err != nil {
				errChan <- err
				continue
			}
			errChan <- h.CloudFileCreate(e, stat)

		case fsnotify.Remove:
			dir, file, isDir := h.GetDeleteID(e)
			if isDir {
				errChan <- h.CloudDirDelete(*dir)
				continue
			}
			errChan <- h.CloudFileDelete(*file)

		case fsnotify.Rename:

			//n(d)] = ed[le

		}

	}

}
