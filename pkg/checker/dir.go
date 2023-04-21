package checker

import (
	"fmt"
	"fsm_client/pkg/database"
	"fsm_client/pkg/ent"
	"fsm_client/pkg/ignore"
	"log"
	"os"
	"path/filepath"
	"time"
)

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

	//log.Println("local client diradd", dirAdd)
	//log.Println("local client dirdelete", dirDelete)

	for _, dir := range dirDelete {
		err := c.Client.DirDelete(dir)
		if err != nil {
			log.Println(err)
		}

		c.DB.Delete(&dir)
	}

	for _, dir := range dirAdd {
		if err := c.Client.DirCreate(&dir); err != nil {
			return err
		}

		c.DB.Create(&dir)
	}

	ch := make(chan int)

	go func() {
		for k := range syncTask {
			cloud, err := c.Client.GetAllDirBySyncID(k)
			if err = c.insertDirs(cloudDir, cloud); err != nil {
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

	//log.Println("cloud  diradd", dirAdd)
	//log.Println("cloud dirdelete", dirDelete)

	for _, dir := range dirDelete {
		dirPath := filepath.Join(syncTask[dir.SyncID], dir.Dir)
		ignore.Lock.Store(dirPath, 1)

		if err := os.RemoveAll(dirPath); err != nil {
			return err
		}
		c.DB.Delete(&dir)

		time.Sleep(time.Millisecond * 400)
		ignore.Lock.Delete(dirPath)
	}

	for _, dir := range dirAdd {
		dirPath := filepath.Join(syncTask[dir.SyncID], dir.Dir)
		ignore.Lock.Store(dirPath, 1)

		if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
			return err
		}
		c.DB.Create(&dir)

		time.Sleep(time.Millisecond * 400)
		ignore.Lock.Delete(dirPath)
	}

	return nil
}
