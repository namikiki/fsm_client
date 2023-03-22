package handle

import (
	"errors"
	"io"
	"log"
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

	_, err = io.Copy(createFile, fileIO)

	fileIO.Close()
	createFile.Close()

	if err := os.Chtimes(filepath.Join(rootPath, parentPath, file.Name), time.Unix(file.ModTime, 0), time.Unix(file.ModTime, 0)); err != nil {
		return err
	}

	h.DB.Save(&file)
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
		CreateTime: time.Now().Unix(),
		ModTime:    stat.ModTime().Unix(),
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

	time.Sleep(time.Millisecond * 500)

	h.DB.Where("sync_id = ? and level =? and name =?", fw.SyncID, level, stat.Name()).Find(&files)
	switch len(files) {
	case 0:
		return h.CloudFileCreate(fw, stat)
	case 1:
		file = &files[0]
	default:
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
		CreateTime: time.Now().Unix(),
		ModTime:    stat.ModTime().Unix(),
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

func (h *Handle) GetRenameData(stat os.FileInfo, event fsn.FsEventWithID, a *[3]string) (file *ent.Dir, dir *ent.File) {

	if stat == nil {
		log.Println("该名之前的路径", event.Path())
		log.Println("等到改名之前的的夫路径", filepath.Dir(event.Path()))
		a[0] = filepath.Dir(event.Path())
		return h.GetDeleteID(event)
	}

	log.Println("该名之后的路径", event.Path())
	a[1] = filepath.Dir(event.Path())

	if stat.IsDir() {
		a[2] = event.AbsPath + PathSeparator
	} else {
		a[2] = filepath.Base(event.Path())
	}

	return nil, nil
}

func (h *Handle) Rename(eventChan chan fsn.FsEventWithID) {
	var a [3]string
	var file *ent.File
	var dir *ent.Dir

	for {

		event := <-eventChan
		time.Sleep(time.Millisecond * 500)
		stat, _ := os.Stat(event.Path())
		d, f := h.GetRenameData(stat, event, &a)

		if d != nil || f != nil {
			file = f
			dir = d
		}

		log.Println("第一次检测")

		timer := time.NewTimer(time.Second * 2)
		select {
		case event2 := <-eventChan:
			info, _ := os.Stat(event2.Path())
			dd, ff := h.GetRenameData(info, event2, &a)

			if dd != nil || ff != nil {
				file = ff
				dir = dd
			}

			log.Println("第二次检测")
		case <-timer.C:
			continue
		}

		//log.Println(a[0])
		//log.Println(a[1])
		//log.Println(a[2])
		//log.Println(file)
		//log.Println(dir)

		if a[0] == a[1] {
			if file != nil {
				file.Name = a[2]
				h.HttpClient.FileRename(*file)
				h.DB.Save(file)
			}
			if dir != nil {
				dir.Dir = a[2]
				h.HttpClient.DirRename(*dir)
				h.DB.Save(dir)
			}
		}

		file, dir = nil, nil
	}
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
			//case notify.Rename:
			//	 h.Rename(e)
		}
	}

}
