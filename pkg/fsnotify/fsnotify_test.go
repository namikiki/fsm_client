package fsn

import (
	"log"
	"path/filepath"
	"testing"
	"time"

	"github.com/rjeczalik/notify"
)

func TestName(t *testing.T) {
	c := make(chan notify.EventInfo, 100)
	err := notify.Watch(filepath.Join("C:\\Users\\surflabom\\Desktop\\MyDediServer", "..."), c, notify.All)
	if err != nil {
		log.Println(err)
		return
	}

	for {
		e := <-c
		log.Println(e.Path(), e.Event())
	}

}

// import (
//
//	"log"
//	"testing"
//
//	"fsm_client/pkg/config"
//	"fsm_client/pkg/ignore"
//
//	"github.com/fsnotify/fsnotify"
//	"github.com/rjeczalik/notify"
//
// )
//
// func TestNew(t *testing.T) {
//
//		watchers := make(map[string]Watch)
//		watcher, err := NewWatch("123", "/Users/zylzyl/go/src/fsm_client/pkg/mock/test")
//		if err != nil {
//			log.Println(err)
//		}
//		watchers[watcher.Path] = watcher
//
//		newIgnoreConfig, err := config.NewIgnoreConfig()
//		if err != nil {
//			log.Println(err)
//		}
//
//		newIgnore, err := ignore.NewIgnore(newIgnoreConfig)
//		if err != nil {
//			log.Println(err)
//		}
//
//		buffChannal := make(chan fsnotify.Event)
//		errBuffChannal := make(chan error)
//		var s chan interface{}
//		manger := NewWatchManger(watchers, newIgnore, buffChannal, errBuffChannal, s)
//		go manger.Watch()
//
//		for {
//			select {
//			case e := <-manger.BuffChannel:
//				log.Println(e.Op, e.Name, e.String())
//			case err := <-manger.ErrBuffChannel:
//				log.Println(err)
//			}
//		}
//	}
func TestOTH(t *testing.T) {
	c := make(chan notify.EventInfo, 1)
	if err := notify.Watch("/Users/zylzyl/go/src/fsm_client/pkg/mock/test/...", c, notify.Create, notify.Remove); err != nil {
		log.Fatal(err)
	}
	defer notify.Stop(c)

	for {
		select {
		case event := <-c:
			log.Println(event.Event().String(), event.Path())
		}
	}
}

func TestName2(t *testing.T) {

	dad := make(chan int, 2)
	timer := time.NewTimer(time.Second * 2)
	dad <- 1

	select {
	case i := <-dad:
		log.Println("s")
		time.Sleep(time.Second * 4)
		log.Println(i)
	case t := <-timer.C:
		log.Println(t)
	}

}
