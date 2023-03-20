package sync

import (
	"log"
	"path/filepath"
	"strings"
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
	//stat, err := os.Stat("/Users/zylzyl/go/src/fsm_client/pkg/mock/gggg")
	//if err != nil {
	//	log.Println(err)
	//}
	//log.Println(stat)

	log.Println(filepath.Join("/Users/zylzyl/go/src/fsm_client/pkg/mock", "..."))

}

func BenchmarkName(b *testing.B) {

	s := "/root/dir1/sub1"

	for i := 0; i < b.N; i++ {
		log.Println(strings.HasPrefix(s, "/root/dir1"))
	}

}
