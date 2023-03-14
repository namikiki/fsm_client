package fsn

import (
	"log"

	"fsm_client/pkg/ignore"

	"github.com/fsnotify/fsnotify"
)

type watch struct {
	W    *fsnotify.Watcher
	ID   string
	Path string
}
type WatchManger struct {
	Watchers       map[string]watch
	Ignore         *ignore.Ignore
	BuffChannel    chan interface{}
	ErrBuffChannel chan error
}

func NewWatchManger(watcher map[string]watch, ignore *ignore.Ignore, buffChannel chan interface{}, errBuffChannel chan error) *WatchManger {
	return &WatchManger{
		Watchers:       watcher,
		Ignore:         ignore,
		BuffChannel:    buffChannel,
		ErrBuffChannel: errBuffChannel,
	}
}

func (wm *WatchManger) Add() {

}

func (wm *WatchManger) Remove() {

}

func (wm *WatchManger) Watch() {
	for s := range wm.Watchers {

		go func() {
			wm.Watchers[s].W.Add(s)
			for {
				select {
				case event := <-wm.Watchers[s].W.Events:
					wm.BuffChannel <- event
				case err := <-wm.Watchers[s].W.Errors:
					wm.ErrBuffChannel <- err
				}
			}
		}()

	}
}

func (wm *WatchManger) ProessChan() {
	for {
		select {
		case event := <-wm.BuffChannel:
			f := event.(*fsnotify.Event)
			if wm.Ignore.Match(f.String()) {
				continue
			}
			log.Println(f.Name)
		case err := <-wm.ErrBuffChannel:
			log.Println(err)
		}
	}
}

type Task struct {
	Name string
	Op   string
}

var TasksChan = make(chan Task, 1)

func StartTask() {
	tasks := map[string]chan int{}

	for {
		select {
		case t := <-TasksChan:
			if t.Op == "add" {
				c := make(chan int)
				go NewFSN(t.Name, c)
				tasks[t.Name] = c

			} else if t.Op == "del" {
				tasks[t.Name] <- 1
				delete(tasks, t.Name)
			}
		}
	}
}

func NewFSN(path string, c chan int) {
	// Create new watcher.
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	// Start listening for events.
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println(path+"event:", event)
				if event.Has(fsnotify.Write) {
					log.Println("modified file:", event.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	// Add a path.
	err = watcher.Add("/tmp/" + path)
	if err != nil {
		log.Fatal(err)
	}

	// Block main goroutine forever.
	<-c
}

func NewWatch() *fsnotify.Watcher {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	return watcher
}

func StartFSN(watcher *fsnotify.Watcher) {

	for {
		select {
		case event := <-watcher.Events:
			log.Println("event:", event)
			if event.Has(fsnotify.Write) {
				log.Println("modified file:", event.Name)
			}
		case err := <-watcher.Errors:
			log.Println("error:", err)
		}
	}

}
