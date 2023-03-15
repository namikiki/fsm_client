package fsn

import (
	"log"
	"os"

	"fsm_client/pkg/ignore"

	"github.com/fsnotify/fsnotify"
)

type Watch struct {
	W      *fsnotify.Watcher
	SyncID string
	Path   string
}

func NewWatch(syncID, path string) (Watch, error) {
	_, err := os.Stat(path)
	if err != nil {
		return Watch{}, err
	}

	watcher, err := fsnotify.NewWatcher()

	return Watch{
		SyncID: syncID,
		Path:   path,
		W:      watcher,
	}, err

}

type WatchManger struct {
	Watchers       map[string]Watch
	Ignore         *ignore.Ignore
	BuffChannel    chan fsnotify.Event
	ErrBuffChannel chan error
}

func NewWatchManger(watchers map[string]Watch, ignore *ignore.Ignore, buffChannel chan fsnotify.Event, errBuffChannel chan error) *WatchManger {
	return &WatchManger{
		Watchers:       watchers,
		Ignore:         ignore,
		BuffChannel:    buffChannel,
		ErrBuffChannel: errBuffChannel,
	}
}

func (wm *WatchManger) Add() {

}

func (wm *WatchManger) Remove() {
	//fsnotify.Th
}

func (wm *WatchManger) Watch() {
	log.Println("watch booting")
	for s := range wm.Watchers {

		go func() {
			wm.Watchers[s].W.Add(wm.Watchers[s].Path)
			for {
				log.Println("watch thread ...")
				select {
				case event := <-wm.Watchers[s].W.Events:
					wm.BuffChannel <- event
				case err := <-wm.Watchers[s].W.Errors:
					wm.ErrBuffChannel <- err
				}
			}
		}()

	}
	log.Println("watch start")
	select {}
}

func (wm *WatchManger) ProessChan() {
	log.Println("proess watch start")
	for {
		select {
		case event := <-wm.BuffChannel:
			if wm.Ignore.Match(event.String()) {
				continue
			}
			log.Println(event.Name)

		case err := <-wm.ErrBuffChannel:
			log.Println(err)
		}
	}
}

//type Task struct {
//	Name string
//	Op   string
//}
//
//var TasksChan = make(chan Task, 1)
//
//func StartTask() {
//	tasks := map[string]chan int{}
//
//	for {
//		select {
//		case t := <-TasksChan:
//			if t.Op == "add" {
//				c := make(chan int)
//				go NewFSN(t.Name, c)
//				tasks[t.Name] = c
//
//			} else if t.Op == "del" {
//				tasks[t.Name] <- 1
//				delete(tasks, t.Name)
//			}
//		}
//	}
//}
//
//func NewFSN(path string, c chan int) {
//	// Create new watcher.
//	watcher, err := fsnotify.NewWatcher()
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer watcher.Close()
//
//	// Start listening for events.
//	go func() {
//		for {
//			select {
//			case event, ok := <-watcher.Events:
//				if !ok {
//					return
//				}
//				log.Println(path+"event:", event)
//				if event.Has(fsnotify.Write) {
//					log.Println("modified file:", event.Name)
//				}
//			case err, ok := <-watcher.Errors:
//				if !ok {
//					return
//				}
//				log.Println("error:", err)
//			}
//		}
//	}()
//
//	// Add a path.
//	err = watcher.Add("/tmp/" + path)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Block main goroutine forever.
//	<-c
//}
//
//func NewWatch() *fsnotify.Watcher {
//	watcher, err := fsnotify.NewWatcher()
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	return watcher
//}
//
//func StartFSN(watcher *fsnotify.Watcher) {
//
//	for {
//		select {
//		case event := <-watcher.Events:
//			log.Println("event:", event)
//			if event.Has(fsnotify.Write) {
//				log.Println("modified file:", event.Name)
//			}
//		case err := <-watcher.Errors:
//			log.Println("error:", err)
//		}
//	}
//
//}
