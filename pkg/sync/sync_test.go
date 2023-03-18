package sync

import (
	"log"
	"os"
	"testing"
)

func TestT1(t *testing.T) {

	//if err != nil {
	//	panic(err)
	//}
	//
	//Ignore, err := ignore.NewIgnore(IgnoreConfig)
	//if err != nil {
	//	log.Println(err)
	//}
	//buffChannal := make(chan fsnotify.Event, 100)
	//
	//manger := fsn.NewWatchManger(buffChannal, Ignore)
	//go manger.Watch()
	//
	//log.Println("文件变化监视器已启动....")
	//
	////user
	//regis := mock.NewRegis()

}

func TestName(t *testing.T) {
	stat, err := os.Stat("/Users/zylzyl/go/src/fsm_client/pkg/mock/gggg")
	if err != nil {
		log.Println(err)
	}
	log.Println(stat)

}
