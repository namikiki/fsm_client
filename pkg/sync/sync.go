package sync

import (
	"encoding/json"
	"errors"
	"log"
	"path/filepath"
	"strings"
	"time"

	"fsm_client/pkg/ent"
	"fsm_client/pkg/handle"
	"fsm_client/pkg/httpclient"
	"fsm_client/pkg/types"

	"fsm_client/pkg/fsnotify"

	"gorm.io/gorm"
)

type Syncer struct {
	httpClient  *httpclient.Client //http client
	DB          *gorm.DB
	Handle      *handle.Handle
	SyncTask    map[string]string
	WatchManger *fsn.WatchManger
}

func NewSyncer(client *httpclient.Client, db *gorm.DB, handle *handle.Handle, watchManger *fsn.WatchManger) *Syncer {
	synctask := make(map[string]string)

	return &Syncer{
		httpClient:  client,
		DB:          db,
		Handle:      handle,
		SyncTask:    synctask,
		WatchManger: watchManger,
	}
}

func (s *Syncer) ListenLocalChanges() error {
	<-s.httpClient.Ch

	var st []ent.SyncTask
	s.DB.Where("status = ?", "sync").Find(&st)

	for _, task := range st {
		watch, err := fsn.NewWatch(task.ID, task.RootDir, task.Ignore)
		if err != nil {
			log.Println(err)
			continue
		}
		s.WatchManger.AddNotifyChannel <- watch
		s.SyncTask[task.ID] = task.RootDir
	}

	for i := 0; i < 3; i++ {
		go s.Handle.PressLocalChange(s.WatchManger.EventWithIDChan, s.WatchManger.ErrBuffChannel)
	}
	s.Handle.Rename(s.WatchManger.RenameChannel)

	return nil
}

func (s *Syncer) ListenCloudDataChanges() error {
	<-s.httpClient.Ch
	log.Println("开启云端文件变化消息获取")

	connect, err := s.httpClient.WebSocketConnect()
	if err != nil {
		panic(err)
	}

	for {
		_, receivedMessage, err := connect.ReadMessage()
		if err != nil {
			log.Fatal("接收消息失败：", err)
		}

		var psm types.PubSubMessage
		if err := json.Unmarshal(receivedMessage, &psm); err != nil {
			log.Println(err)
		}

		if psm.ClientID == s.httpClient.ClientID {
			continue
		}

		if _, ok := s.SyncTask[psm.SyncID]; !ok && psm.Type != "syncTask" {
			log.Println("未知消息，屏蔽", psm.Type, psm.Action)
			continue
		}

		switch psm.Type {
		case "file":

			var file ent.File
			var dir ent.Dir
			json.Unmarshal(psm.Data, &file)

			s.DB.Where("id = ?", file.ParentDirID).Find(&dir)
			s.Handle.FileChange(psm.Action, file, dir.Dir, s.SyncTask[file.SyncID])

		case "dir":

			var dir ent.Dir
			json.Unmarshal(psm.Data, &dir)
			s.Handle.DirChange(psm.Action, dir, s.SyncTask[dir.SyncID])

		case "syncTask":
			var synctask ent.SyncTask
			json.Unmarshal(psm.Data, &synctask)

			if psm.Action == "create" {
				synctask.Status = "created"
				s.Handle.SyncTaskCreate(synctask)

			} else {
				synctask.Status = "delete"
				s.WatchManger.RemoveNotifyChannel <- synctask.ID
				delete(s.SyncTask, synctask.ID)
			}

		default:
			log.Println("未知事件", psm.Type, psm.Action, psm.SyncID)
		}

	}
}

func (s *Syncer) Error() {
	for {
		if err := <-s.WatchManger.ErrBuffChannel; err != nil {
			log.Println(err)
		}
	}
}

func (s *Syncer) CreateSyncTask(st types.NewSyncTask) error {

	task := ent.SyncTask{
		Type:       st.Type,
		Name:       st.Name,
		RootDir:    st.Path,
		Ignore:     st.Ignore,
		CreateTime: time.Now().Unix(),
	}

	var gs []ent.SyncTask
	s.DB.Find(&gs)
	for _, st := range gs {
		if prefix := strings.HasPrefix(task.RootDir, st.RootDir); prefix {
			return errors.New("不能添加子目录")
		}
		// todo add "/" or "\"
		if prefix := strings.HasPrefix(st.RootDir, task.RootDir); prefix {
			return errors.New("不能添加父目录")
		}
	}

	//todo 添加任务时，应该在全量备份完成后 通知其他客户端添加了同步任务
	task.RootDir = filepath.ToSlash(st.Path)
	if err := s.httpClient.SyncTaskCreate(&task); err != nil {
		return err
	}

	task.RootDir = st.Path
	task.Status = "syncing"
	s.DB.Create(&task)

	if err := s.Handle.ScannerPathToUpload(task.RootDir, task.ID); err != nil {
		return err
	}

	watch, err := fsn.NewWatch(task.ID, task.RootDir, task.Ignore)
	if err != nil {
		return err
	}

	s.SyncTask[task.ID] = task.RootDir
	s.WatchManger.AddNotifyChannel <- watch

	task.Status = "sync"
	s.DB.Save(&task)
	return err
}

func (s *Syncer) RecoverTask(st types.RecSyncTask) error {
	var synctask ent.SyncTask

	if s.DB.Where("id = ?", st.ID).Find(&synctask); synctask.Status != "created" {
		return errors.New("该任务不可被恢复同步状态")
	}

	synctask.Name = st.Name
	synctask.RootDir = st.Path
	synctask.Ignore = st.Ignore
	synctask.Status = "syncing"
	s.DB.Save(&synctask)

	err := s.Handle.GetSyncTaskToDownload(synctask.ID, synctask.RootDir)
	if err != nil {
		return err
	}

	watch, err := fsn.NewWatch(synctask.ID, synctask.RootDir, synctask.Ignore)
	if err != nil {
		return err
	}
	s.WatchManger.AddNotifyChannel <- watch

	s.SyncTask[synctask.ID] = synctask.RootDir
	synctask.Status = "sync"
	s.DB.Save(&synctask)

	return err
}

// PauseAndContinueTask todo  增加暂停时长
func (s *Syncer) PauseAndContinueTask(syncID string) error {
	var task ent.SyncTask
	if s.DB.Where("id = ?", syncID).Find(&task); task.ID == "" {
		return errors.New("未找到")
	}

	if task.Status == "pause" {
		watch, err := fsn.NewWatch(task.ID, task.RootDir, task.Ignore)
		if err != nil {
			return err
		}

		log.Println("重新同步")
		task.Status = "sync"
		s.SyncTask[task.ID] = task.RootDir
		s.WatchManger.AddNotifyChannel <- watch
	} else if task.Status == "sync" || task.Status == "syncing" {
		task.Status = "pause"
		s.WatchManger.RemoveNotifyChannel <- syncID
		delete(s.SyncTask, syncID)
	}

	s.DB.Save(&task)
	return nil
}

func (s *Syncer) DeleteSyncTask(del types.DeleteSyncTask) error {

	var sync ent.SyncTask
	if s.DB.Where("id = ?", del.ID).Find(&sync); sync.ID == "" {
		return errors.New("未找到")
	}

	//todo 添加延迟 保证在监视者被完全删除 或者 让 RemoveNotifyChannel 返回信号
	s.WatchManger.RemoveNotifyChannel <- sync.ID
	delete(s.SyncTask, sync.ID)

	if del.DelCloud {
		s.DB.Delete(&sync)
		if err := s.httpClient.SyncTaskDelete(del.ID); err != nil {
			return err
		}
	}

	sync.Status = "delete"
	s.DB.Save(&sync)

	if del.DelLocal {
		return s.Handle.DeleteAllFileByDir(sync.RootDir)
	}

	return nil
}

//func (s *Syncer) CancelSyncTask(syncID string) error {
//	var sync ent.SyncTask
//	s.DB.Where("id=?", syncID).Find(&sync)
//
//	//todo 关闭本地文件夹监控
//	//todo 将数据库 更新为 delete
//	//todo 过滤掉云端的数据变化消息
//	return nil
//}
