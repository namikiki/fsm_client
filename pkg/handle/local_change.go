package handle

import (
	"errors"
	"fsm_client/pkg/ent"
	fsn "fsm_client/pkg/fsnotify"
	"github.com/rjeczalik/notify"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func (h *Handle) PressLocalChange(eventChan chan fsn.FsEventWithID, errChan chan error) {
	for {
		e := <-eventChan
		switch e.Event {
		case notify.Create:
			stat, err := os.Stat(e.Path)
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
				errChan <- errors.New("未找到删除的文件或者文件夹的位置" + e.Path)
			}
		}
	}

}

//file

func (h *Handle) CloudFileCreate(fw fsn.FsEventWithID, stat os.FileInfo) error {
	level := len(strings.Split(fw.AbsPath, PathSeparator)) - 1

	var dir ent.Dir
	suffix := strings.TrimSuffix(fw.AbsPath, "/"+stat.Name())
	h.DB.Where("dir = ? and level = ?", suffix, level).Find(&dir)

	file := ent.File{
		SyncID:      fw.SyncID,
		Name:        stat.Name(),
		Level:       level,
		Deleted:     false,
		CreateTime:  time.Now().Unix(),
		ModTime:     stat.ModTime().Unix(),
		ParentDirID: dir.ID,
	}

	fileIO, err := os.Open(fw.Path)
	if err != nil {
		return err
	}

	if err := h.HttpClient.FileCreate(&file, fileIO); err != nil {
		return err
	}

	h.DB.Create(&file)
	return nil
}

func (h *Handle) GetDeleteID(fw fsn.FsEventWithID) (*ent.Dir, *ent.File) {
	level := len(strings.Split(fw.AbsPath, PathSeparator)) - 1
	fileName := filepath.Base(fw.AbsPath)

	log.Println(fw.SyncID, level, fileName)
	var files []ent.File
	h.DB.Where("sync_id = ? and level =? and name =?", fw.SyncID, level, fileName).Find(&files)

	switch len(files) {
	case 0:
		log.Println("寻找被删除的文件夹")
		var adir ent.Dir
		if h.DB.Where("sync_id = ? and level =? and dir =?", fw.SyncID, level+1, fw.AbsPath).Find(&adir); adir.ID != "" {
			return &adir, nil
		}
	case 1:
		return nil, &files[0]
	default:

		var dir ent.Dir
		var file ent.File
		suffix := strings.TrimSuffix(fw.AbsPath, "/"+fileName)
		h.DB.Where("dir = ? and level = ?", suffix, level).Find(&dir)

		log.Println(dir)
		log.Println(fw.SyncID, level, dir.ID, fileName)

		h.DB.Where("sync_id = ? and level =? and parent_dir_id= ? and name =?", fw.SyncID, level, dir.ID, fileName).Find(&file)
		return nil, &file
	}

	return nil, nil
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

func (h *Handle) CloudFileDelete(file ent.File) error {
	if err := h.HttpClient.FileDelete(file); err != nil {
		return err
	}

	h.DB.Delete(&file)
	return nil
}

func (h *Handle) CloudFileUpdate(fw fsn.FsEventWithID) error {
	level := len(strings.Split(fw.AbsPath, PathSeparator)) - 1

	var file *ent.File

	time.Sleep(time.Millisecond * 500)

	stat, err := os.Stat(fw.Path)
	if err != nil {
		return err
	}

	var dir ent.Dir
	suffix := strings.TrimSuffix(fw.AbsPath, "/"+stat.Name())
	h.DB.Where("dir = ? and level = ?", suffix, level).Find(&dir)

	log.Println(dir)
	log.Println(fw.SyncID, level, dir.ID, stat.Name())

	h.DB.Where("sync_id = ? and level =? and parent_dir_id= ? and name =?", fw.SyncID, level, dir.ID, stat.Name()).Find(&file)
	log.Println(file)

	if file.ID == "" {
		return h.CloudFileCreate(fw, stat)
	}

	fileIO, err := os.Open(fw.Path)
	if err != nil {
		return err
	}

	file.ModTime = stat.ModTime().Unix()
	if err = h.HttpClient.FileUpdate(file, fileIO); err != nil {
		return err
	}

	h.DB.Save(file)
	return nil
}

//dir

func (h *Handle) CloudDirCreate(fw fsn.FsEventWithID, stat os.FileInfo) error {
	level := len(strings.Split(fw.AbsPath, PathSeparator))

	dir := ent.Dir{
		SyncID:     fw.SyncID,
		Dir:        fw.AbsPath,
		Level:      level,
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

//rename

func (h *Handle) GetRenameData(stat os.FileInfo, event fsn.FsEventWithID, a *[3]string) (file *ent.Dir, dir *ent.File) {
	if stat == nil {
		log.Println("该名之前的路径", event.Path)
		log.Println("等到改名之前的的夫路径", filepath.Dir(event.Path))
		a[0] = filepath.Dir(event.Path)
		log.Println("该名之前的路径aaaaaaa", event.Path)
		return h.GetDeleteID(event)
	}

	log.Println("该名之后的路径", event.Path)
	a[1] = filepath.Dir(event.Path)

	if stat.IsDir() {
		a[2] = event.AbsPath
	} else {
		a[2] = filepath.Base(event.Path)
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
		stat, _ := os.Stat(event.Path)
		d, f := h.GetRenameData(stat, event, &a)

		if d != nil || f != nil {
			file = f
			dir = d
		}

		log.Println("第一次检测")

		timer := time.NewTimer(time.Second * 2)
		select {
		case event2 := <-eventChan:
			info, _ := os.Stat(event2.Path)
			dd, ff := h.GetRenameData(info, event2, &a)

			if dd != nil || ff != nil {
				file = ff
				dir = dd
			}

			log.Println("第二次检测")
		case <-timer.C:
			continue
		}

		log.Println(a[0])
		log.Println(a[1])
		log.Println(a[2])
		log.Println(file)
		log.Println(dir)

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
