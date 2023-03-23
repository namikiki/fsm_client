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

			mock.NewLogin,
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
		),
	).Err()

	if err != nil {
		log.Println(err)
	}

	select {}
}
