package sync

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"path/filepath"
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

func (s *Syncer) ListenCloudDataChanges() error {
	connect, err := s.httpClient.WebSocketConnect()
	if err != nil {
		panic(err)
	}

	for {
		messageType, receivedMessage, err := connect.ReadMessage()
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
			switch psm.Action {
			case "create":
				var file ent.File
				json.Unmarshal(psm.Data, &file)
				fileIO, err := s.httpClient.GetFile(file.ID)
				if err != nil {
					return err
				}
				var dir ent.Dir
				s.DB.Where("id = ?", file.ParentDirID).Find(&dir)
				f, err := os.Create(filepath.Join(dir.Dir, file.Name))
				if err != nil {
					return err
				}

				io.Copy(f, fileIO)
				f.Close()
				fileIO.Close()

			case "delete":
			case "update":

			}
		case "dir":
			var dir ent.Dir
			json.Unmarshal(psm.Data, &dir)
			if psm.Action == "create" {
				s.Handle.DirCreate(psm.Data, s.SyncTask[dir.SyncID])
				continue
			}
			s.Handle.DirDelete(psm.Data, s.SyncTask[dir.SyncID])

		case "synctask":
			if psm.Action == "create" {
				s.Handle.SyncTaskCreate(psm.Data)
				continue
			}
			s.Handle.SyncTaskDelete(psm.Data)
		default:
			log.Println(psm)
		}

		log.Printf("接收到的消息类型：%d \n msg %v", messageType, psm)

		//fmt.Printf()
		//fmt.Printf("接收到的消息内容：%s\n", string(receivedMessage))
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
