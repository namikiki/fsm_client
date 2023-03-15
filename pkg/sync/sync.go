package sync

import (
	"encoding/json"
	"log"
	"time"

	"fsm_client/pkg/ent"
	"fsm_client/pkg/handle"
	"fsm_client/pkg/httpclient"
	"fsm_client/pkg/types"

	"gorm.io/gorm"
)

type Syncer struct {
	httpClient *httpclient.Client //http client
	DB         *gorm.DB
	Handle     *handle.Handle
	SyncTask   map[string]string
}

func NewSyncer(client *httpclient.Client, db *gorm.DB, handle *handle.Handle) *Syncer {
	// todo  init SyncTask
	synctask := make(map[string]string)

	return &Syncer{
		httpClient: client,
		DB:         db,
		Handle:     handle,
		SyncTask:   synctask,
	}
}

func (s *Syncer) TaskInit() {
	var st []ent.SyncTask
	s.DB.Find(&st)
	for _, task := range st {
		s.SyncTask[task.ID] = task.RootDir
	}
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

	return s.Handle.ScannerPathToUpload(task.RootDir, task.ID)
}

func (s *Syncer) RestoreSyncTask(taskID, path string) error {
	//todo 更改同步任务状态
	err := s.Handle.GetSyncTaskToDownload(taskID, path)
	if err != nil {
		return err
	}

	//todo 开启数据变化DIR监视
	return err
}

func (s *Syncer) PauseSyncTask() {

}

func (s *Syncer) ContinuePause() {

}

//func (s *Syncer) GetFile(fileID string) (io.ReadCloser, error) {
//	request, _ := http.NewRequest("GET", baseUrl+"/file/open/"+fileID, nil)
//	request.Header.Set("Content-Type", "application/json")
//
//	resp, err := s.httpClient.Do(request)
//	return resp.Body, err
//}
//
//func (s *Syncer) GetAllDirBySyncID(syncID string) ([]ent.Dir, error) {
//	request, _ := http.NewRequest("GET", baseUrl+"/dir/getAllDirBySyncID/"+syncID, nil)
//	request.Header.Set("Content-Type", "application/json")
//
//	var dirs []ent.Dir
//	if resp, err := s.httpClient.Do(request); err == nil {
//
//		return dirs, json.NewDecoder(resp.Body).Decode(&dirs)
//	}
//	return nil, nil
//}
//
//func (s *Syncer) GetAllFileBySyncID(syncID string) ([]ent.File, error) {
//
//	request, _ := http.NewRequest("GET", baseUrl+"/file/get/all/bySyncID/"+syncID, nil)
//	request.Header.Set("Content-Type", "application/json")
//
//	var files []ent.File
//	if resp, err := s.httpClient.Do(request); err == nil {
//		return files, json.NewDecoder(resp.Body).Decode(&files)
//	}
//	return nil, nil
//}
//
//func (s *Syncer) CreateSyncTask(task *ent.SyncTask) error {
//	marshal, _ := json.Marshal(task)
//	request, _ := http.NewRequest("POST", baseUrl+"/synctask/create", bytes.NewBuffer(marshal))
//	request.Header.Set("Content-Type", "application/json")
//	request.Header.Set("client", s.clientID)
//	if resp, err := s.httpClient.Do(request); err == nil {
//		return json.NewDecoder(resp.Body).Decode(&task)
//	}
//	return nil
//}
//
//func (s *Syncer) CreateDir(e *ent.Dir) error {
//	return nil
//}
//
//func (s *Syncer) CreateFile(e *ent.File, open *os.File) error {
//	return nil
//}
