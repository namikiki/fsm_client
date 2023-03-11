package sync

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"fsm_client/pkg/ent"

	"github.com/google/go-querystring/query"
	"gorm.io/gorm"
)

const base_url = "http://127.0.0.1:8080"

type Syncer struct {
	httpClient *http.Client //http client
	db         *gorm.DB
	clientID   string
}

func NewSyncer(client *http.Client, db *gorm.DB, clientID string) *Syncer {
	return &Syncer{
		httpClient: client,
		db:         db,
		clientID:   clientID,
	}
}

func (s *Syncer) CreateDir(dir *ent.Dir) error {
	marshal, _ := json.Marshal(dir)

	request, _ := http.NewRequest("POST", base_url+"/dir/create", bytes.NewBuffer(marshal))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("client", s.clientID)

	if resp, err := s.httpClient.Do(request); err == nil {
		return json.NewDecoder(resp.Body).Decode(&dir)
	}
	return nil
}

func (s *Syncer) CreateFile(file *ent.File, fileio io.ReadCloser) error {
	defer fileio.Close()
	values, _ := query.Values(file)

	request, _ := http.NewRequest("POST", base_url+"/file/create?"+values.Encode(), fileio)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("client", s.clientID)

	if resp, err := s.httpClient.Do(request); err == nil {
		return json.NewDecoder(resp.Body).Decode(file)
	}
	return nil
}

func (s *Syncer) GetFile(fileID string) (io.ReadCloser, error) {
	request, _ := http.NewRequest("GET", base_url+"/file/open/"+fileID, nil)
	request.Header.Set("Content-Type", "application/json")

	resp, err := s.httpClient.Do(request)
	return resp.Body, err
}

func (s *Syncer) GetAllDirBySyncID(syncID string) ([]ent.Dir, error) {
	request, _ := http.NewRequest("GET", base_url+"/dir/getAllDirBySyncID/"+syncID, nil)
	request.Header.Set("Content-Type", "application/json")

	var dirs []ent.Dir
	if resp, err := s.httpClient.Do(request); err == nil {

		return dirs, json.NewDecoder(resp.Body).Decode(&dirs)
	}
	return nil, nil
}

func (s *Syncer) GetAllFileBySyncID(syncID string) ([]ent.File, error) {

	request, _ := http.NewRequest("GET", base_url+"/file/get/all/bySyncID/"+syncID, nil)
	request.Header.Set("Content-Type", "application/json")

	var files []ent.File
	if resp, err := s.httpClient.Do(request); err == nil {
		return files, json.NewDecoder(resp.Body).Decode(&files)
	}
	return nil, nil
}

func (s *Syncer) CreateSyncTask(task *ent.SyncTask) error {
	marshal, _ := json.Marshal(task)
	request, _ := http.NewRequest("POST", base_url+"/synctask/create", bytes.NewBuffer(marshal))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("client", s.clientID)
	if resp, err := s.httpClient.Do(request); err == nil {
		return json.NewDecoder(resp.Body).Decode(&task)
	}
	return nil
}
