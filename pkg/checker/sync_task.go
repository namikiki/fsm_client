package checker

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"fsm_client/pkg/database"
	"fsm_client/pkg/ent"
)

func (c *Checker) GetSyncTaskChange() map[string]string {
	var task []ent.SyncTask
	c.DB.Find(&task)

	localSyncDelete := make(map[string]string)
	localSync := make(map[string]string)
	cloudSync := make(map[string]string)

	for _, t := range task {
		localSync[t.ID] = t.RootDir
		_, err := os.Stat(t.RootDir)
		if err != nil {
			localSyncDelete[t.ID] = t.RootDir
		}
	}

	syncTasks, err := c.Client.SyncTaskGetAll()
	if err != nil {
		log.Println(err)
	}
	for _, s := range syncTasks {
		cloudSync[s.ID] = s.RootDir
	}

	cloudDelete := make(map[string]string)
	cloudAdd := make(map[string]string)

	for key := range localSync {
		if _, ok := cloudSync[key]; !ok {
			cloudDelete[key] = localSync[key]
		}
	}

	for key := range cloudSync {
		if _, ok := localSync[key]; !ok {
			cloudAdd[key] = cloudSync[key]
		}
	}

	log.Println("add")
	for i, i2 := range cloudAdd {
		log.Println(i, i2)
	}

	log.Println("delete")
	for i, i2 := range cloudDelete {
		log.Println(i, i2)
	}

	return localSync
}

func (c *Checker) GetDriveFileAndDir(syncTasks map[string]string) {

	for syncID, rootDir := range syncTasks {
		rootPathLen := len(rootDir)

		filepath.WalkDir(rootDir, func(path string, d fs.DirEntry, err error) error {
			path = path[rootPathLen:]

			if d.IsDir() {
				dir := ent.Dir{
					SyncID:     syncID,
					Dir:        path + "/", // todo Split
					Level:      uint64(len(strings.Split(path, "/"))),
					Deleted:    false,
					CreateTime: time.Now().Unix(),
					ModTime:    time.Now().Unix(),
				}

				c.insertDir(DriveDir, dir)
				return err
			}

			level := len(strings.Split(path, "/"))
			suffix := strings.TrimSuffix(path, d.Name())
			//var dir ent.Dir
			//c.MemDB.Where("dir = ? and level = ?", suffix, level-1).Find(&dir)

			info, _ := d.Info()
			file := ent.File{
				SyncID:      syncID,
				Name:        d.Name(),
				ParentDirID: suffix,
				Level:       uint64(level),
				Size:        info.Size(),
				Deleted:     false,
				CreateTime:  time.Now().Unix(),
				ModTime:     info.ModTime().Unix(),
			}

			c.insertFile(DriveFile, file)
			return err
		})

	}

}

func (c *Checker) GetDirChange(syncTask map[string]string) error {
	var dirs []ent.Dir
	c.DB.Find(&dirs)
	c.insertDirs(DBDir, dirs)

	var dirDelete, dirAdd []ent.Dir

	GetDirAddSQL := fmt.Sprintf(GetDirChange, DBDir, DBDir, DriveDir, DBDir, DriveDir, DBDir, DriveDir, DriveDir)
	if err := c.MemDB.Select(&dirDelete, GetDirAddSQL); err != nil {
		return err
	}

	GetDirDeleteSQL := fmt.Sprintf(GetDirChange, DriveDir, DriveDir, DBDir, DriveDir, DBDir, DriveDir, DBDir, DBDir)
	if err := c.MemDB.Select(&dirAdd, GetDirDeleteSQL); err != nil {
		return err
	}

	log.Println("local client diradd", dirAdd)
	log.Println("local client dirdelete", dirDelete)

	//todo 屏蔽
	for _, dir := range dirDelete {

		err := c.Client.DirDelete(dir)
		if err != nil {
			log.Println(err)
		}

		if err := os.RemoveAll(filepath.Join(syncTask[dir.SyncID], dir.Dir)); err != nil {
			return err
		}
		c.DB.Delete(&dir)
	}

	for _, dir := range dirAdd {

		if err := c.Client.DirCreate(&dir); err != nil {
			return err
		}

		if err := os.MkdirAll(filepath.Join(syncTask[dir.SyncID], dir.Dir), os.ModePerm); err != nil {
			return err
		}
		c.DB.Create(&dir)
	}

	ch := make(chan int)

	go func() {
		for k := range syncTask {
			cloud, err := c.Client.GetAllDirBySyncID(k)
			err = c.insertDirs(cloudDir, cloud)
			if err != nil {
				log.Println(err)
			}
		}

		ch <- 1
	}()

	database.ResetDirTable(c.MemDB)
	c.DB.Find(&dirs)
	c.insertDirs(DBDir, dirs)

	<-ch

	GetCloudDirAddSQL := fmt.Sprintf(GetDirChange, DBDir, DBDir, cloudDir, DBDir, cloudDir, DBDir, cloudDir, cloudDir)
	if err := c.MemDB.Select(&dirDelete, GetCloudDirAddSQL); err != nil {
		return err
	}

	GetCloudDirDeleteSQL := fmt.Sprintf(GetDirChange, cloudDir, cloudDir, DBDir, cloudDir, DBDir, cloudDir, DBDir, DBDir)
	if err := c.MemDB.Select(&dirAdd, GetCloudDirDeleteSQL); err != nil {
		return err
	}

	log.Println("cloud  diradd", dirAdd)
	log.Println("cloud dirdelete", dirDelete)

	//todo 屏蔽
	for _, dir := range dirDelete {

		err := c.Client.DirDelete(dir)
		if err != nil {
			log.Println(err)
		}

		if err := os.RemoveAll(filepath.Join(syncTask[dir.SyncID], dir.Dir)); err != nil {
			return err
		}
		c.DB.Delete(&dir)
	}

	for _, dir := range dirAdd {

		if err := c.Client.DirCreate(&dir); err != nil {
			return err
		}

		if err := os.MkdirAll(filepath.Join(syncTask[dir.SyncID], dir.Dir), os.ModePerm); err != nil {
			return err
		}
		c.DB.Create(&dir)
	}

	return nil
}

func (c *Checker) GetFileChange(syncTask map[string]string) error {
	var files []ent.File
	c.DB.Find(&files)
	c.insertFiles(DBFile, files)

	var fileAdd, fileDelete, fileUpdate []ent.File
	if err := c.MemDB.Select(&fileDelete, GetFileChange(DBFile, DriveFile)); err != nil {
		log.Println(err)
		return err
	}

	if err := c.MemDB.Select(&fileAdd, GetFileChange(DriveFile, DBFile)); err != nil {
		log.Println(err)
		return err
	}

	if err := c.MemDB.Select(&fileUpdate, getFileUpdate(DBFile, DriveFile)); err != nil {
		log.Println(err)
		return err
	}

	log.Println("local file add", fileAdd)
	log.Println("local file delete", fileDelete)
	log.Println("local file update", fileUpdate)

	for _, file := range fileAdd {
		var dir ent.Dir
		c.DB.Where("sync_id = ? and dir = ? ", file.SyncID, file.ParentDirID).Find(&dir)
		file.ParentDirID = dir.ID
		fileIO, err := os.Open(syncTask[file.SyncID] + dir.Dir + file.Name)
		if err != nil {
			log.Println(err)
			return err
		}

		err = c.Client.FileCreate(&file, fileIO)
		if err != nil {
			return err
		}
		c.DB.Create(&file)
	}

	for _, file := range fileDelete {
		if err := c.Client.FileDelete(file); err != nil {
			return err
		}
		c.DB.Delete(&file)
	}

	for _, file := range fileUpdate {
		var dir ent.Dir
		c.DB.Where("sync_id = ? and id = ? ", file.SyncID, file.ParentDirID).Find(&dir)

		fileIO, err := os.Open(syncTask[file.SyncID] + dir.Dir + file.Name)

		stat, err := fileIO.Stat()
		file.Size = stat.Size()
		file.ModTime = stat.ModTime().Unix()

		if err != nil {
			log.Println(err)
			return err
		}

		err = c.Client.FileUpdate(&file, fileIO)
		if err != nil {
			log.Println(err)
			return err
		}
		log.Println(file)

		c.DB.Save(&file)
	}

	ch := make(chan int)

	go func() {
		for k := range syncTask {
			cloud, err := c.Client.GetAllFileBySyncID(k)
			err = c.insertFiles(cloudFile, cloud)
			if err != nil {
				log.Println(err)
			}
		}

		ch <- 1
	}()

	database.ResetFileTable(c.MemDB)
	c.DB.Find(&files)
	c.insertFiles(DBFile, files)

	<-ch

	if err := c.MemDB.Select(&fileDelete, GetFileChange(DBFile, cloudFile)); err != nil {
		log.Println(err)
		return err
	}

	if err := c.MemDB.Select(&fileAdd, GetFileChange(cloudFile, DBFile)); err != nil {
		log.Println(err)
		return err
	}

	if err := c.MemDB.Select(&fileUpdate, getFileUpdate(cloudFile, DBFile)); err != nil {
		log.Println(err)
		return err
	}

	log.Println("cloud file add", fileAdd)
	log.Println("cloud file delete", fileDelete)
	log.Println("cloud file update", fileUpdate)

	for _, file := range fileAdd {
		var dir ent.Dir
		c.DB.Where("sync_id = ? and id = ? ", file.SyncID, file.ParentDirID).Find(&dir)

		err := c.Handle.FileWrite(file, dir.Dir, syncTask[file.SyncID])
		if err != nil {
			log.Println(err)
			return err
		}

	}

	for _, file := range fileDelete {

		var dir ent.Dir
		c.DB.Where("sync_id = ? and id = ? ", file.SyncID, file.ParentDirID).Find(&dir)
		err := c.Handle.FileDelete(file, dir.Dir, syncTask[file.SyncID])

		if err != nil {
			log.Println(err)
			return err
		}

		if err := c.Client.FileDelete(file); err != nil {
			return err
		}
		c.DB.Delete(&file)
	}

	for _, file := range fileUpdate {
		var dir ent.Dir
		c.DB.Where("sync_id = ? and id = ? ", file.SyncID, file.ParentDirID).Find(&dir)

		err := c.Handle.FileWrite(file, dir.Dir, syncTask[file.SyncID])
		if err != nil {
			log.Println(err)
			return err
		}
	}

	return nil
}
