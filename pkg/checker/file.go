package checker

import (
	"fsm_client/pkg/database"
	"fsm_client/pkg/ent"
	"fsm_client/pkg/ignore"
	"log"
	"os"
	"path/filepath"
	"time"
)

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
		if dir.ID == "" {
			log.Println("未找到添加文件的文件夹")
			return nil
		}

		file.ParentDirID = dir.ID
		fileIO, err := os.Open(filepath.Join(syncTask[file.SyncID], dir.Dir, file.Name))
		if err != nil {
			log.Println(err)
			return err
		}

		if err = c.Client.FileCreate(&file, fileIO); err != nil {
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

		fileIO, err := os.Open(filepath.Join(syncTask[file.SyncID], dir.Dir, file.Name))

		stat, err := fileIO.Stat()
		if err != nil {
			log.Println(err)
			return err
		}
		file.ModTime = stat.ModTime().Unix()

		if err = c.Client.FileUpdate(&file, fileIO); err != nil {
			log.Println(err)
			return err
		}

		c.DB.Save(&file)
	}

	ch := make(chan int)

	go func() {
		for k := range syncTask {
			cloud, err := c.Client.GetAllFileBySyncID(k)
			if err = c.insertFiles(cloudFile, cloud); err != nil {
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

		filePath := filepath.Join(syncTask[file.SyncID], dir.Dir, file.Name)
		ignore.Lock.Store(filePath, 1)

		if err := c.Handle.FileWrite(file, filePath); err != nil {
			log.Println(err)
			return err
		}

		time.Sleep(time.Millisecond * 400)
		ignore.Lock.Delete(filePath)
	}

	for _, file := range fileDelete {

		var dir ent.Dir
		c.DB.Where("sync_id = ? and id = ? ", file.SyncID, file.ParentDirID).Find(&dir)

		filePath := filepath.Join(syncTask[file.SyncID], dir.Dir, file.Name)
		ignore.Lock.Store(filePath, 1)

		if err := c.Handle.FileDelete(file, filePath); err != nil {
			log.Println(err)
			return err
		}

		time.Sleep(time.Millisecond * 400)
		ignore.Lock.Delete(filePath)
	}

	for _, file := range fileUpdate {
		var dir ent.Dir
		c.DB.Where("sync_id = ? and id = ? ", file.SyncID, file.ParentDirID).Find(&dir)

		filePath := filepath.Join(syncTask[file.SyncID], dir.Dir, file.Name)
		ignore.Lock.Store(filePath, 1)

		if err := c.Handle.FileWrite(file, filePath); err != nil {
			log.Println(err)
			return err
		}

		time.Sleep(time.Millisecond * 400)
		ignore.Lock.Delete(filePath)
	}

	return nil
}
