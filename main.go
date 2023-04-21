package main

import (
	"fmt"
	httpapi "fsm_client/api/http"
	"fsm_client/pkg/checker"
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
	"github.com/google/uuid"
	"go.uber.org/fx"
	"gorm.io/gorm"
	"log"
	"os"
	"os/signal"
	"syscall"
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

			checker.NewChecker,
			database.NewSqliteMemoryDB,

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

			//
			//func(client *httpclient.Client, log types.UserLoginReq) {
			//	client.Login(log)
			//},

			func(client *httpclient.Client, check *checker.Checker) {
				log.Println("jwt login start")
				if err := client.LoginByJWT(); err != nil {
					log.Println(err)
				}
				log.Println("jwt login end")

				go func() {
					log.Println("check wait")
					<-client.Ch
					log.Println("check start")
					checkTask := check.GetSyncTaskChange()
					log.Println("1111111111111")
					if err := check.GetDriveFileAndDir(checkTask); err != nil {
						log.Println(err)
						return
					}

					log.Println("22222222222222")
					if err := check.GetDirChange(checkTask); err != nil {
						log.Println(err)
						return
					}

					log.Println("33333333333333")
					if err := check.GetFileChange(checkTask); err != nil {
						log.Println(err)
						return
					}
				}()

				log.Println("check over")
			},

			//func() {
			//
			//},

			//func(sync *sync.Syncer) {
			//	st := types.NewSyncTask{
			//		Name:   "test",
			//		Path:   "C:\\Users\\surflabom\\Desktop\\lib",
			//		Type:   "two",
			//		Ignore: false,
			//	}
			//	//
			//	if err := sync.CreateSyncTask(st); err != nil {
			//		log.Println(err)
			//		return
			//	}
			//},

			//
			//	log.Println("停止同步任务")
			//
			//	var syncc ent.SyncTask
			//	if sync.DB.Where("name = ?", "test").Find(&syncc); syncc.ID == "" {
			//		log.Println("未找到")
			//		return
			//	}
			//
			//	if err := sync.PauseAndContinueTask(syncc.ID); err != nil {
			//		log.Println(err)
			//		return
			//	}
			//
			//	time.Sleep(time.Second * 3)
			//
			//	if err := sync.PauseAndContinueTask(syncc.ID); err != nil {
			//		log.Println(err)
			//		return
			//	}
			//},

			//func(sync *sync.Syncer) {
			//	log.Println("停止同步任务....")
			//	time.Sleep(time.Second * 120)
			//
			//
			//},

			//func(sync *sync.Syncer) {
			//	st := types.RecSyncTask{
			//		ID:     "d000cb22-4eb8-4504-8a86-358eb1fdf86d",
			//		Name:   "test",
			//		Path:   "C:\\Users\\surflabom\\Desktop\\rec",
			//		Ignore: false,
			//	}
			//
			//	err := sync.RecoverTask(st)
			//	if err != nil {
			//		return
			//	}
			//},
			//
		),
	).Err()

	if err != nil {
		log.Println(err)
	}

	go func() {
		if err := api.App.Run("127.0.0.1:9000"); err != nil {
			panic(err)
			log.Fatal(err)
		}
	}()

	// 创建一个通道来接收信号
	sigs := make(chan os.Signal, 1)

	// 监听 SIGINT 和 SIGTERM 信号
	signal.Notify(sigs, syscall.SIGKILL, syscall.SIGTERM)

	// 等待接收信号
	sig := <-sigs
	fmt.Println("接收到信号：", sig)
}
