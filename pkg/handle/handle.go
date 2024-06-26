package handle

import (
	"errors"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"runtime"
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
	isWindows := runtime.GOOS == "windows"

	rootPathLen := len(rooPath)
	return filepath.WalkDir(rooPath, func(path string, d fs.DirEntry, err error) error {

		absPath := path[rootPathLen:]
		if isWindows {
			absPath = filepath.ToSlash(absPath)
		}

		if d.IsDir() {

			dir := ent.Dir{
				SyncID:     syncID,
				Dir:        absPath,
				Level:      len(strings.Split(absPath, "/")),
				Deleted:    false,
				CreateTime: time.Now().Unix(),
				ModTime:    time.Now().Unix(),
			}

			if err := h.HttpClient.DirCreate(&dir); err != nil {
				return err
			}

			h.DB.Create(&dir)
			return err
		}

		var dir ent.Dir
		info, _ := d.Info()

		level := len(strings.Split(absPath, "/")) - 1
		suffix := strings.TrimSuffix(absPath, "/"+info.Name())

		if h.DB.Where("dir= ? and level = ?", suffix, level).Find(&dir); dir.ID == "" {
			return errors.New(" not found dir")
		}

		file := ent.File{
			SyncID:      syncID,
			Name:        d.Name(),
			ParentDirID: dir.ID,
			Level:       level,
			Deleted:     false,
			CreateTime:  time.Now().Unix(),
			ModTime:     info.ModTime().Unix(),
		}

		fileIO, err := os.Open(path)
		if err := h.HttpClient.FileCreate(&file, fileIO); err != nil {
			return err
		}

		h.DB.Create(&file)
		return err
	})

}

func (h *Handle) GetSyncTaskToDownload(syncID, path string) error {
	dirs, err := h.HttpClient.GetAllDirBySyncID(syncID)
	if err != nil {
		return err
	}

	for i := range dirs {
		if err := os.MkdirAll(path+dirs[i].Dir, os.ModePerm); err != nil {
			if !os.IsExist(err) {
				log.Println(err)
				return err
			}
		} // 文件夹创建成功
	}
	h.DB.Create(&dirs)

	files, err := h.HttpClient.GetAllFileBySyncID(syncID)
	if err != nil {
		return err
	}

	for i := range files {

		fileIO, err := h.HttpClient.GetFile(files[i].ID)
		if err != nil {
			return err
		}

		var dir ent.Dir
		h.DB.Where("id = ?", files[i].ParentDirID).Find(&dir)

		filePath := filepath.Join(path, dir.Dir, files[i].Name)
		file, err := os.Create(filePath)
		if err != nil {
			return err
		}

		if _, err = io.Copy(file, fileIO); err != nil {
			return err
		}

		file.Close()
		fileIO.Close()

		os.Chtimes(filePath, time.Now(), time.Unix(files[i].ModTime, 0))
	} //文件创建成功
	h.DB.Create(&files)

	return nil
}
