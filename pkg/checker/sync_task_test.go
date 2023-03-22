package checker

import (
	"log"
	"testing"

	"fsm_client/pkg/config"
	"fsm_client/pkg/database"
	"fsm_client/pkg/handle"
	"fsm_client/pkg/httpclient"
	"fsm_client/pkg/mock"

	"github.com/google/uuid"
)

func initCheck() *Checker {

	cfg, _ := config.ReadConfigFile()
	cfg.Device.ClientID = uuid.NewString()
	log.Println(cfg)

	memoryDB := database.NewSqliteMemoryDB()
	db := database.NewGormSQLiteConnect()
	Client := httpclient.NewClient(cfg)

	regis := mock.NewRegis()

	handle := handle.NewHandle(Client, db, nil)

	err := Client.Register(regis)
	if err != nil {
		log.Println(err)
	}

	login := mock.NewLogin()
	err = Client.Login(login)
	if err != nil {
		log.Println(err)
	}

	checker := NewChecker(memoryDB, db, Client, handle, nil)
	return checker
}

func TestRT1(t *testing.T) {
	checker := initCheck()
	syncTaskChange := checker.GetSyncTaskChange()
	log.Println(syncTaskChange)
	checker.GetDriveFileAndDir(syncTaskChange)
	checker.GetDirChange(syncTaskChange)
	checker.GetFileChange(syncTaskChange)
}

func TestT2(t *testing.T) {
	//checker := initCheck()
}

//func TestName(t *testing.T) {
//
//	local := map[string]int{"a": 1, "b": 2, "c": 3}
//	clound := map[string]int{"a": 1, "b": 3, "d": 6}
//
//	cloundDelete := make(map[string]int)
//	cloundAdd := make(map[string]int)
//
//	// 检查map1中是否有map2没有的键
//	for key := range local {
//		if _, ok := clound[key]; !ok {
//			cloundDelete[key] = local[key]
//		}
//	}
//
//	// 检查map2中是否有map1没有的键
//	for key := range clound {
//		if _, ok := local[key]; !ok {
//			cloundDelete[key] = clound[key]
//		}
//	}
//
//
//
//	for k, v := range cloundDelete {
//		log.Println(k, v)
//	}
//
//	log.Println("------")
//
//	for k, v := range cloundAdd {
//		log.Println(k, v)
//	}
//
//
//	noma tsak map[syncid]path
//
//	for i, i2 := range collection {
//
//		filepath.WalkDir("", func(path string, d fs.DirEntry, err error) error {
//
//			c.db.create
//
//		})
//
//
//
//
//		getDirbySyncid
//
//		getFileBySyncid
//
//	}
//
//
//
//}
