package main

import (
	"log"
	"os"

	"fsm_client/pkg/config"
	"fsm_client/pkg/database"
	fsn "fsm_client/pkg/fsnotify"
	"fsm_client/pkg/handle"
	"fsm_client/pkg/httpclient"
	"fsm_client/pkg/ignore"
	"fsm_client/pkg/mock"
	"fsm_client/pkg/sync"
	"fsm_client/pkg/types"

	"github.com/google/uuid"
	"go.uber.org/fx"
)

func Init() (int64, *types.Config) {
	cfg, _ := config.ReadConfigFile()
	cfg.Device.ClientID = uuid.NewString()
	log.Println(cfg)
	return 100, cfg
}

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)

	err := fx.New(

		fx.Provide(
			Init,
			config.NewIgnoreConfig,
			ignore.NewIgnore,
			fsn.NewWatchManger,

			mock.NewRegis,
			httpclient.NewClient,

			mock.NewAccount,
			database.NewGormSQLiteConnect,
			handle.NewHandle,
			sync.NewSyncer,
		),

		fx.Invoke(
			func(client *httpclient.Client, reg types.UserRegister) {
				client.Register(reg)
			},

			func(client *httpclient.Client, account types.UserLoginReq) {
				client.Login(account)
			},

			func(sync *sync.Syncer) {
				if sync.WatchManger == nil {
					log.Println("error")
				}

				go sync.WatchManger.Watch()
				go sync.Error()
				go sync.ListenLocalChanges()
			},

			func(sync *sync.Syncer) {
				go sync.ListenCloudDataChanges()
			},

			func(sync *sync.Syncer) {
				if os.Args[1] != "sla" {
					sync.CreateSyncTask("test", "/Users/zylzyl/go/src/fsm_client/test/client1/src")
				}
			},

			//func(sync *sync.Syncer) {
			//	sync.RestoreSyncTask("bda5ffca-947b-4456-93a5-043f91466273", "/Users/zylzyl/Desktop/markdown/synctest/tyu")
			//},
		),
	).Err()
	if err != nil {
		log.Println(err)
	}

	select {}

	//app.Start(context.Background())
	//log.Println("233")
	//c := make(chan notify.EventInfo, 10)
	//if err := notify.Watch(filepath.Join("/Users/zylzyl/go/src/fsm_client/pkg/mock", "..."), c, notify.All); err != nil {
	//	log.Fatal(err)
	//}
	//defer notify.Stop(c)
	//
	//for {
	//	select {
	//	case event := <-c:
	//		ei := event.Sys().(*notify.FSEvent)
	//		log.Println(ei.ID, ei.Path, ei.Flags)
	//		log.Println(event.Event().String(), event.Sys(), event.Path())
	//	}
	//}

	//log.SetFlags(log.LstdFlags | log.Llongfile)
	//sla()
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

	//gormCon := database.NewGormSQLiteConnect()
	//hand := handle.NewHandle(client, gormCon)
	//syncer := sync.NewSyncer(client, gormCon, hand, nil)

	//go func() {
	//	err := syncer.ListenCloudDataChanges()
	//	if err != nil {
	//		log.Println(err)
	//	}
	//}()

	//err := syncer.CreateSyncTask("syncTest", "/Users/zylzyl/Desktop/GolangProjects/fsm/filetest")
	//if err != nil {
	//	log.Println(err)
	//}
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
	//
	//gormCon := database.NewGormSQLiteConnect()
	//hand := handle.NewHandle(client, gormCon)
	//syncer := sync.NewSyncer(client, gormCon, hand, nil)
	//
	//syncer.ListenLocalChanges()
	//
	//go func() {
	//	err := syncer.ListenCloudDataChanges()
	//	if err != nil {
	//		log.Println(err)
	//	}
	//}()
	//
	//select {}
	//
	//err := syncer.CreateSyncTask("syncTest", "/Users/zylzyl/Desktop/GolangProjects/fsm/filetest")
	//if err != nil {
	//	log.Println(err)
	//}
}
