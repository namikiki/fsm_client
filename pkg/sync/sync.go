package sync

import (
	"encoding/json"
	"log"
	"time"

	//"time"

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
	// todo  init SyncTask
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
	var st []ent.SyncTask
	s.DB.Find(&st)

	for _, task := range st {
		watch, err := fsn.NewWatch(task.ID, task.RootDir, true)
		if err != nil {
			return err
		}
		s.WatchManger.AddChannel <- watch
	}

	s.Handle.PressLocalChange(s.WatchManger.EventWithIDChan, s.WatchManger.ErrBuffChannel)
	return nil
}

func (s *Syncer) ListenCloudDataChanges() error {
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

		switch psm.Type {
		case "file":
			var file ent.File
			var dir ent.Dir
			json.Unmarshal(psm.Data, &file)
			s.DB.Where("id = ?", file.ParentDirID).Find(&dir)

			if psm.Action == "update" || psm.Action == "create" {
				s.Handle.FileWrite(file, dir.Dir, s.SyncTask[file.SyncID])
			} else {
				s.Handle.FileDelete(file, dir.Dir, s.SyncTask[file.SyncID])
			}

		case "dir":

			var dir ent.Dir
			json.Unmarshal(psm.Data, &dir)
			if psm.Action == "create" {
				s.Handle.DirCreate(dir, s.SyncTask[dir.SyncID])
			} else {
				s.Handle.DirDelete(dir, s.SyncTask[dir.SyncID])
			}

		case "syncTask":
			var synctask ent.SyncTask
			json.Unmarshal(psm.Data, &synctask)

			if psm.Action == "create" {
				s.Handle.SyncTaskCreate(synctask)
			} else {
				s.Handle.SyncTaskDelete(synctask)
			}

		default:
			log.Println(psm)
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

func (s *Syncer) CreateSyncTask(name, root string) error {

	task := ent.SyncTask{
		Type:       "two",
		Name:       name,
		RootDir:    root,
		Deleted:    false,
		CreateTime: time.Now(),
	}

	if err := s.httpClient.SyncTaskCreate(&task); err != nil {
		return err
	}
	s.DB.Create(&task)

	if err := s.Handle.ScannerPathToUpload(task.RootDir, task.ID); err != nil {
		return err
	}

	watch, err := fsn.NewWatch(task.ID, task.RootDir, true)
	if err != nil {
		return err
	}

	s.WatchManger.AddChannel <- watch
	return err
}

func (s *Syncer) DeleteSyncTask(syncID string, deleteFile bool) error {
	var sync ent.SyncTask
	s.DB.Where("id=?", syncID).Find(&sync)

	if deleteFile {
		return s.Handle.DeleteAllFileByDir(sync.RootDir)
	}
	return s.httpClient.SyncTaskDelete(syncID)
}

func (s *Syncer) CancelSyncTask(syncID string) error {
	var sync ent.SyncTask
	s.DB.Where("id=?", syncID).Find(&sync)

	//todo 关闭本地文件夹监控
	//todo 将数据库 更新为 delete
	//todo 过滤掉云端的数据变化消息
	return nil
}

func (s *Syncer) PauseSyncTask() {

	//syncID
	//path
	//bool
	// ignore
	//
	//
	//s.WatchManger.AddAndDeleteChannel <-
	// todo 关闭本地文件夹监控
	//todo 过滤掉云端的数据变化消息
}

func (s *Syncer) Continue() {

}

func (s *Syncer) RestoreSyncTask(taskID, path string) error {
	//todo 更改同步任务状态
	err := s.Handle.GetSyncTaskToDownload(taskID, path)
	if err != nil {
		return err
	}

	watch, err := fsn.NewWatch(taskID, path, true)
	if err != nil {
		return err
	}
	s.WatchManger.AddChannel <- watch
	return err
}
