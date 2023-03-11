package main

import (
	"log"

	"fsm_client/pkg/config"
	"fsm_client/pkg/http"
	"fsm_client/pkg/mock"
)

func main() {

	cfg, _ := config.ReadConfigFile()
	client := http.NewClient("http://127.0.0.1:8080", "ws://127.0.0.1:8080ß", cfg)

	account := mock.NewAccount()
	if err := client.Login(account); err != nil {
		log.Println(err)
	}

	dir := mock.NewDir()
	if err := client.DirCreate(&dir); err != nil {
		log.Println(err)
	}

	if err := client.DirDelete(dir); err != nil {
		log.Println(err)
	}

	log.Println(dir)

	//go syncer.WebSocketConn()
	//
	//err := syncer.GetSyncTaskToDownload("6fd35b63-95fb-46b8-a7fa-394ccec20b01", "/Users/zylzyl/Desktop/GolangProjects/fsm/test/filetest")
	//if err != nil {
	//	log.Println(err)
	//}

	//err := syncer.CreateSync("filetest", "/Users/zylzyl/Desktop/GolangProjects/fsm/test/filetest")
	//if err != nil {
	//	log.Println(err)
	//}
	//
	//select {}
}
