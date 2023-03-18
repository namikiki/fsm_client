package fsn

import (
	"log"
	"os"

	"fsm_client/pkg/ent"
	"fsm_client/pkg/ignore"

	"github.com/fsnotify/fsnotify"
)

type Watch struct {
	W      *fsnotify.Watcher
	SyncID string
	Path   string
	Ignore bool
}

type FsEventWithID struct {
	fsnotify.Event
	SyncID  string
	AbsPath string
}

type FileFsEvent struct {
	Event    fsnotify.Event
	FullPath string
	AbsPath  string
	File     ent.File
}

type DirFsEvent struct {
	Event fsnotify.Event
	Dir   ent.Dir
}

func NewWatch(syncID, path string, ignore bool) (Watch, error) {
	_, err := os.Stat(path)
	if err != nil {
		return Watch{}, err
	}

	watcher, err := fsnotify.NewWatcher()
	return Watch{
		SyncID: syncID,
		Path:   path,
		W:      watcher,
		Ignore: ignore,
	}, err

}

type WatchManger struct {
	Watchers        map[string]Watch
	EventWithIDChan chan FsEventWithID
	ErrBuffChannel  chan error
	AddChannel      chan Watch
	RemoveChannel   chan string
	Ignore          *ignore.Ignore
}

func NewWatchManger(buffLen int64,
	ignore *ignore.Ignore) *WatchManger {

	eventWithIDChan := make(chan FsEventWithID, buffLen)
	errBuffChannel := make(chan error, 2)
	removeChannel := make(chan string, 2)
	addChannel := make(chan Watch, 4)
	watchers := make(map[string]Watch)

	return &WatchManger{
		Watchers:        watchers,
		AddChannel:      addChannel,
		RemoveChannel:   removeChannel,
		EventWithIDChan: eventWithIDChan,
		ErrBuffChannel:  errBuffChannel,
		Ignore:          ignore,
	}
}

func (wm *WatchManger) add(w Watch) {
	if err := w.W.Add(w.Path); err != nil {
		wm.ErrBuffChannel <- err
		return
	}

	log.Printf("ID %s 路径 %s 开始监控", w.SyncID, w.Path)
	rootPathLen := len(w.Path)
	//PathSeparator := string(os.PathSeparator)

	if w.Ignore {
		for {
			select {
			case event := <-w.W.Events:
				if wm.Ignore.Match(event.Name) {
					continue
				}
				path := event.Name[rootPathLen:]
				wm.EventWithIDChan <- FsEventWithID{event, w.SyncID, path}

			case err := <-w.W.Errors:
				wm.ErrBuffChannel <- err
			}
		}
	}

	for { //非过滤
		select {
		case event := <-w.W.Events:
			path := event.Name[rootPathLen:]
			wm.EventWithIDChan <- FsEventWithID{event, w.SyncID, path}
		case err := <-w.W.Errors:
			wm.ErrBuffChannel <- err
		}
	}

}

//func (wm *WatchManger) add(w Watch) {
//	if err := w.W.Add(w.Path); err != nil {
//		wm.ErrBuffChannel <- err
//	}
//	log.Printf("ID %s 路径 %s 开始监控", w.SyncID, w.Path)
//	rootPathLen := len(w.Path)
//	PathSeparator := string(os.PathSeparator)
//
//	if w.Ignore {
//		for {
//			select {
//			case event := <-w.W.Events:
//				if wm.Ignore.Match(event.Name) {
//					continue
//				}
//
//				stat, err := os.Stat(event.Name)
//				if err != nil {
//					log.Println(err)
//					wm.ErrBuffChannel <- err
//					continue
//				}
//
//				path := event.Name[rootPathLen:]
//				level := len(strings.Split(path, PathSeparator))
//
//				if stat.IsDir() {
//					wm.DirChannel <- DirFsEvent{
//						Event: event,
//						Dir: ent.Dir{
//							SyncID:     w.SyncID,
//							Dir:        path,
//							Level:      uint64(level),
//							Deleted:    false,
//							CreateTime: time.Now(),
//							ModTime:    stat.ModTime(),
//						},
//					}
//
//				} else {
//
//					wm.FileChannel <- FileFsEvent{
//						Event:    event,
//						FullPath: event.Name,
//						AbsPath:  path,
//						File: ent.File{
//							SyncID:      w.SyncID,
//							Name:        stat.Name(),
//							ParentDirID: "",
//							Level:       uint64(level),
//							Deleted:     false,
//							CreateTime:  time.Now(),
//							ModTime:     stat.ModTime(),
//						},
//					}
//				}
//			case err := <-w.W.Errors:
//				wm.ErrBuffChannel <- err
//			}
//		}
//	} else { // 非过滤
//		for {
//			select {
//			case event := <-w.W.Events:
//				stat, err := os.Stat(event.Name)
//				if err != nil {
//					log.Println(err)
//					wm.ErrBuffChannel <- err
//					continue
//				}
//
//				path := event.Name[rootPathLen:]
//				level := len(strings.Split(path, PathSeparator))
//
//				if stat.IsDir() {
//					wm.DirChannel <- DirFsEvent{
//						Event: event,
//						Dir: ent.Dir{
//							SyncID:     w.SyncID,
//							Dir:        path,
//							Level:      uint64(level),
//							Deleted:    false,
//							CreateTime: time.Now(),
//							ModTime:    stat.ModTime(),
//						},
//					}
//
//				} else {
//
//					wm.FileChannel <- FileFsEvent{
//						Event:    event,
//						FullPath: event.Name,
//						AbsPath:  path,
//						File: ent.File{
//							SyncID:      w.SyncID,
//							Name:        stat.Name(),
//							ParentDirID: "",
//							Level:       uint64(level),
//							Deleted:     false,
//							CreateTime:  time.Now(),
//							ModTime:     stat.ModTime(),
//						},
//					}
//				}
//			case err := <-w.W.Errors:
//				wm.ErrBuffChannel <- err
//			}
//		}
//	}
//}

func (wm *WatchManger) remove(syncID string) {
	delete(wm.Watchers, syncID)

	if err := wm.Watchers[syncID].W.Close(); err != nil {
		wm.ErrBuffChannel <- err
	}
}

func (wm *WatchManger) Watch() {

	for {
		select {
		case a := <-wm.AddChannel:
			go wm.add(a)
			wm.Watchers[a.SyncID] = a
		case r := <-wm.RemoveChannel:
			wm.remove(r)
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
