package sync

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestT1(t *testing.T) {

	//9078e73-1 15 false 1679488180 1679482839
	//1679488180 1679482839}
	//2023-03-22 17:09:33 +0800 +0800 2023-03-21 09:59:25 +0800 +0800
	stat, _ := os.Stat("/Users/zylzyl/go/src/fsm_client/test/client1/src/1.txt")
	log.Println(stat.ModTime().Unix())
	//
	//os.Chtimes("/Users/zylzyl/go/src/fsm_client/test/client1/src/123.go", time.Unix(stat.ModTime().Unix(), 0), time.Unix(stat.ModTime().Unix(), 0))
	//
	//s, _ := os.Stat("/Users/zylzyl/go/src/fsm_client/test/client1/src/123.go")
	//log.Println(s.ModTime().Unix())
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
