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

	"github.com/rjeczalik/notify"
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

	fileIO, err := os.Open(fw.Path())
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
	if err := h.HttpClient.FileDelete(file); err != nil {
		return err
	}

	h.DB.Delete(&file)
	return nil
}

func (h *Handle) CloudFileUpdate(fw fsn.FsEventWithID) error {
	split := strings.Split(fw.AbsPath, PathSeparator)
	level := len(split)

	var files []ent.File
	var file *ent.File
	var err error

	stat, err := os.Stat(fw.Path())
	if err != nil {
		return err
	}

	h.DB.Where("sync_id = ? and level =? and name =?", fw.SyncID, level, stat.Name()).Find(&files)
	if len(files) == 0 {
		return h.CloudFileCreate(fw, stat)
	}

	if len(files) == 1 {
		file = &files[0]
	} else {
		if file, err = h.GetUFile(fw.AbsPath, split[level-1], fw.SyncID, level); err != nil {
			return err
		}
	}

	fileIO, err := os.Open(fw.Path())
	if err != nil {
		return err
	}

	if err = h.HttpClient.FileUpdate(file, fileIO); err != nil {
		return err
	}

	h.DB.Save(file)
	return nil
}

func (h *Handle) CloudDirCreate(fw fsn.FsEventWithID, stat os.FileInfo) error {
	level := len(strings.Split(fw.AbsPath, PathSeparator))

	dir := ent.Dir{
		SyncID:     fw.SyncID,
		Dir:        fw.AbsPath + "/",
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

// GetUFile  当出现多个 syncID level name 相同的文件时，获取 file
func (h *Handle) GetUFile(absPath, name, syncID string, level int) (*ent.File, error) {
	parentPath := strings.TrimSuffix(absPath, name)
	var dir ent.Dir
	h.DB.Where("sync_id = ? and level =? and dir =?", syncID, level-1, parentPath).Find(&dir)
	if dir.ID == "" {
		return nil, errors.New("")
	}

	var file ent.File
	h.DB.Where("sync_id = ? and level =? and parent_dir_id =? and name =?", syncID, level, dir.ID, name).Find(&file)
	if file.ID == "" {
		return nil, errors.New("")
	}
	return &file, nil
}

func (h *Handle) GetDeleteID(fw fsn.FsEventWithID) (*ent.Dir, *ent.File) {
	split := strings.Split(fw.AbsPath, PathSeparator)
	level := len(split)
	var dir ent.Dir

	var files []ent.File
	h.DB.Where("sync_id = ? and level =? and name =?", fw.SyncID, level, split[level-1]).Find(&files)

	switch len(files) {
	case 0:
		if h.DB.Where("sync_id = ? and level =? and dir =?", fw.SyncID, level, fw.AbsPath+"/").Find(&dir); dir.ID != "" {
			return &dir, nil
		}
	case 1:
		return nil, &files[0]
	default:
		if file, err := h.GetUFile(fw.AbsPath, split[level-1], fw.SyncID, level); err != nil {
			return nil, file
		}
		return nil, nil
	}

	return nil, nil
}

func (h *Handle) PressLocalChange(eventChan chan fsn.FsEventWithID, errChan chan error) {
	for {
		e := <-eventChan
		switch e.Event() {
		case notify.Create:
			stat, err := os.Stat(e.Path())
			if err != nil {
				errChan <- err
				continue
			}

			if stat.IsDir() {
				errChan <- h.CloudDirCreate(e, stat)
				continue
			}
			errChan <- h.CloudFileUpdate(e)
			//errChan <- h.CloudFileCreate(e, stat)
		case notify.Write:
			errChan <- h.CloudFileUpdate(e)
		case notify.Remove:
			dir, file := h.GetDeleteID(e)
			if dir != nil {
				errChan <- h.CloudDirDelete(*dir)
			} else if file != nil {
				errChan <- h.CloudFileDelete(*file)
			} else {
				errChan <- errors.New("未找到删除的文件或者文件夹的位置" + e.Path())
			}
		case notify.Rename:
			//var ty [2]fsn.FsEventWithID
			//ty[len(ty)] = e

		}

	}

}
