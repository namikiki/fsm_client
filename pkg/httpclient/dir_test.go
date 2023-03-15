package httpclient

import (
	"log"
	"testing"

	"fsm_client/pkg/config"
	"fsm_client/pkg/mock"
)

func Init() *Client {
	cfg, _ := config.ReadConfigFile()

	client := NewClient(cfg)

	account := mock.NewAccount()
	if err := client.Login(account); err != nil {
		log.Println(err)
	}

	return client
}

func TestCreateDir(t *testing.T) {
	client := Init()
	dir := mock.NewDir()

	if err := client.DirCreate(&dir); err != nil {
		log.Println(err)
	}
	log.Println(dir)
}

func TestDeleteDir(t *testing.T) {
	client := Init()
	dir := mock.NewDir()

	if err := client.DirDelete(dir); err != nil {
		log.Println(err)
	}
}

func TestGetAllDirBySyncID(t *testing.T) {
	client := Init()
	dir := mock.NewDir()

	dirs, err := client.GetAllDirBySyncID(dir.SyncID)
	if err != nil {
		log.Println(err)
	}
	log.Println(dirs)
}
