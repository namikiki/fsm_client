//go:build windows

package handle

import (
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"fsm_client/pkg/ent"
	"fsm_client/pkg/httpclient"
	"fsm_client/pkg/ignore"

	"gorm.io/gorm"
)

type Handle struct {
	HttpClient *httpclient.Client
	DB         *gorm.DB
	Ignore     *ignore.Ignore
}

// NewHandle create a handle
func NewHandle(c *httpclient.Client, db *gorm.DB, ignore *ignore.Ignore) *Handle {
	return &Handle{
		HttpClient: c,
		DB:         db,
		Ignore:     ignore,
	}
}

func (h *Handle) ScannerPathToUpload(rooPath, syncID string) error {

	rootPathLen := len(rooPath)
	return filepath.WalkDir(rooPath, func(path string, d fs.DirEntry, err error) error {

		absPath := filepath.ToSlash(path[rootPathLen:])

		log.Printf("raw= %s, ,abs=%s", path, absPath)

		//if d.IsDir() {
		//
		//	dir := ent.Dir{
		//		SyncID:     syncID,
		//		Dir:        path + "/", // todo Split
		//		Level:      uint64(len(strings.Split(path, "/"))),
		//		Deleted:    false,
		//		CreateTime: time.Now().Unix(),
		//		ModTime:    time.Now().Unix(),
		//	}
		//	if err := h.HttpClient.DirCreate(&dir); err != nil {
		//		return err
		//	}
		//
		//	h.DB.Create(&dir)
		//	return err
		//}

		//level := len(strings.Split(path, "/"))
		//suffix := strings.TrimSuffix(path, d.Name())
		//var dir ent.Dir
		//h.DB.Where("dir= ? and level = ?", suffix, level-1).Find(&dir)
		//info, _ := d.Info()
		//
		//file := ent.File{
		//	SyncID:      syncID,
		//	Name:        d.Name(),
		//	ParentDirID: dir.ID,
		//	Level:       uint64(level),
		//	Size:        info.Size(),
		//	Deleted:     false,
		//	CreateTime:  time.Now().Unix(),
		//	ModTime:     info.ModTime().Unix(),
		//}
		//
		//fileIO, err := os.Open(rawPath)
		//if err := h.HttpClient.FileCreate(&file, fileIO); err != nil {
		//	return err
		//}
		//
		//h.DB.Create(&file)
		return err
	})

}

func (h *Handle) GetSyncTaskToDownload(syncID, path string) error {
	dirs, err := h.HttpClient.GetAllDirBySyncID(syncID)
	if err != nil {
		return err
	}

	for i := range dirs {
		ignore.Lock.Store(strings.TrimSuffix(path+dirs[i].Dir, PathSeparator), 1)

		if err := os.MkdirAll(path+dirs[i].Dir, os.ModePerm); err != nil {
			if !os.IsExist(err) {
				log.Println(err)
				return err
			}
		} // 文件夹创建成功
		time.Sleep(time.Millisecond * 500)
		ignore.Lock.Delete(path + dirs[i].Dir)
	}
	h.DB.Create(&dirs)

	files, err := h.HttpClient.GetAllFileBySyncID(syncID)
	if err != nil {
		return err
	}

	for i := range files {

		if fileIO, err := h.HttpClient.GetFile(files[i].ID); err == nil {
			var dir ent.Dir
			h.DB.Where("id = ?", files[i].ParentDirID).Find(&dir)

			ignore.Lock.Store(path+dir.Dir+files[i].Name, 1)

			if file, err := os.Create(path + dir.Dir + files[i].Name); err == nil {
				io.Copy(file, fileIO)
				file.Close()
			}
			fileIO.Close()

			os.Chtimes(path+dir.Dir+files[i].Name, time.Unix(files[i].ModTime, 0), time.Unix(files[i].ModTime, 0))

			time.Sleep(time.Millisecond * 500)
			ignore.Lock.Delete(path + dir.Dir + files[i].Name)
		}

	} //文件创建成功
	h.DB.Create(&files)
	return nil
}
