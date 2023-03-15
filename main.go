package main

import (
	"log"

	"fsm_client/pkg/config"
	"fsm_client/pkg/database"
	"fsm_client/pkg/handle"
	"fsm_client/pkg/httpclient"
	"fsm_client/pkg/mock"
	"fsm_client/pkg/sync"

	"github.com/google/uuid"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	sla()
	//sla()
	//if os.Args[1] == "mas" {
	//	master()
	//} else {
	//	sla()
	//}

	//err := syncer.RestoreSyncTask("820eb8bd-82b6-4891-aaba-6c93bc96a947", "/Users/zylzyl/go/src/fsm_client/test/filetest")
	//if err != nil {
	//	log.Println(err)
	//}

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

func master() {
	cfg, _ := config.ReadConfigFile()
	log.Println(cfg)

	cfg.Device.ClientID = uuid.NewString()
	client := httpclient.NewClient(cfg)
	regis := mock.NewRegis()
	if err := client.Register(regis); err != nil {
		log.Println(err)
	}

	account := mock.NewAccount()
	if err := client.Login(account); err != nil {
		log.Println(err)
	}

	gormCon := database.NewGormSQLiteConnect()
	hand := handle.NewHandle(client, gormCon)
	syncer := sync.NewSyncer(client, gormCon, hand)

	//go func() {
	//	err := syncer.ListenCloudDataChanges()
	//	if err != nil {
	//		log.Println(err)
	//	}
	//}()

	err := syncer.CreateSyncTask("syncTest", "/Users/zylzyl/Desktop/GolangProjects/fsm/filetest")
	if err != nil {
		log.Println(err)
	}
}

func sla() {

	cfg, _ := config.ReadConfigFile()
	log.Println(cfg)

	cfg.Device.ClientID = uuid.NewString()
	client := httpclient.NewClient(cfg)

	//regis := mock.NewRegis()
	//if err := client.Register(regis); err != nil {
	//	log.Println(err)
	//}

	account := mock.NewAccount()
	if err := client.Login(account); err != nil {
		log.Println(err)
	}

	gormCon := database.NewGormSQLiteConnect()
	hand := handle.NewHandle(client, gormCon)
	syncer := sync.NewSyncer(client, gormCon, hand)

	syncer.TaskInit()

	go func() {
		err := syncer.ListenCloudDataChanges()
		if err != nil {
			log.Println(err)
		}
	}()

	select {}

	err := syncer.CreateSyncTask("syncTest", "/Users/zylzyl/Desktop/GolangProjects/fsm/filetest")
	if err != nil {
		log.Println(err)
	}
}
