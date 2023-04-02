package main

import (
	httpapi "fsm_client/api/http"
	"fsm_client/pkg/config"
	"fsm_client/pkg/database"
	fsn "fsm_client/pkg/fsnotify"
	"fsm_client/pkg/handle"
	"fsm_client/pkg/httpclient"
	"fsm_client/pkg/ignore"
	"fsm_client/pkg/mock"
	"fsm_client/pkg/sync"
	"fsm_client/pkg/types"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"

	"github.com/google/uuid"
	"go.uber.org/fx"
)

func Init() (int64, *types.Config, *gin.Engine) {
	cfg, _ := config.ReadConfigFile()
	cfg.Device.ClientID = uuid.NewString()
	log.Println(cfg)

	engine := gin.Default()

	return 100, cfg, engine
}

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)

	var api *httpapi.Handle

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

			func(e *gin.Engine, client *httpclient.Client, sync *sync.Syncer, db *gorm.DB) {
				api = httpapi.New(e, client, sync, db)
			},

			func(sync *sync.Syncer) {
				go sync.ListenCloudDataChanges()
			},

			func(sync *sync.Syncer) {
				if sync.WatchManger == nil {
					log.Println("error")
				}

				go sync.WatchManger.Watch()
				go sync.Error()
				go sync.ListenLocalChanges()
			},

			func(client *httpclient.Client, reg types.UserRegister) {
				client.Register(reg)
			},

			func(client *httpclient.Client) {
				if err := client.LoginByJWT(); err != nil {
					return
				}
			},

			//func(sync *sync.Syncer) {
			//
			//	st := types.NewSyncTask{
			//		Name:   "test",
			//		Path:   "/Users/zylzyl/go/src/fsm_client/test/client1/src",
			//		Type:   "two",
			//		Ignore: true,
			//	}
			//
			//	if os.Args[1] != "sla" {
			//		sync.CreateSyncTask(st)
			//	}
			//
			//},

		),
	).Err()

	if err != nil {
		log.Println(err)
	}

	if err := api.App.Run(":9000"); err != nil {
		log.Fatal(err)
	}

}
