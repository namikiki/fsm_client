package checker

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"fsm_client/pkg/ent"
)

func (c *Checker) GetSyncTaskChange() map[string]string {
	syncCheckTask := make(map[string]string)
	localSyncMap := make(map[string]ent.SyncTask)
	var localSyncTask []ent.SyncTask

	c.DB.Find(&localSyncTask)

	for _, t := range localSyncTask {
		localSyncMap[t.ID] = t

		if t.Status != "delete" && t.Status != "created" {
			if _, err := os.Stat(t.RootDir); err != nil {
				log.Println("本地删除", t.ID, t.RootDir)
				t.Status = "delete"
				c.DB.Save(&t)
			}
		}

	}

	cloudSyncTask, err := c.Client.SyncTaskGetAll()
	if err != nil {
		log.Println(err)
	}

	for i, s := range cloudSyncTask {
		if _, ok := localSyncMap[s.ID]; !ok {
			log.Println("云端新增", s.ID, s.RootDir)
			cloudSyncTask[i].Status = "created"
			c.DB.Create(&cloudSyncTask[i])
			continue
		}

		delete(localSyncMap, s.ID)
	}

	for _, value := range localSyncMap {
		//本地文件和记录删除
		log.Println("云端删除", value.ID, value.RootDir)
		value.Status = "delete"
		c.DB.Save(&value)
	}

	c.DB.Where("status IN ?", []string{"sync", "syncing", "update"}).Find(&localSyncTask)
	for _, t := range localSyncTask {
		log.Println("本地正常同步", t.ID, t.RootDir)
		syncCheckTask[t.ID] = t.RootDir
	}

	return syncCheckTask
}

func (c *Checker) GetDriveFileAndDir(syncTasks map[string]string) error {
	isWindows := runtime.GOOS == "windows"

	for syncID, rootDir := range syncTasks {

		log.Println("扫描本地文件夹", rootDir)
		rootPathLen := len(rootDir)

		if err := filepath.WalkDir(rootDir, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

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

				return c.insertDir(DriveDir, dir)
			}

			info, _ := d.Info()
			level := len(strings.Split(absPath, "/")) - 1
			suffix := strings.TrimSuffix(absPath, "/"+info.Name())

			//if h.DB.Where("dir= ? and level = ?", suffix, level).Find(&dir); dir.ID == "" {
			//	return errors.New(" not found dir")
			//}

			file := ent.File{
				SyncID:      syncID,
				Name:        d.Name(),
				ParentDirID: suffix,
				Level:       level,
				Size:        info.Size(),
				Deleted:     false,
				CreateTime:  time.Now().Unix(),
				ModTime:     info.ModTime().Unix(),
			}

			//fileIO, err := os.Open(path)
			//if err := h.HttpClient.FileCreate(&file, fileIO); err != nil {
			//	return err
			//}
			c.insertFile(DriveFile, file)
			return err

		}); err != nil {
			return err
		}
	}

	return nil
}
