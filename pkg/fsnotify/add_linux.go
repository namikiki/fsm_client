//go:build linux

package fsn

import (
	"fsm_client/pkg/ignore"
	"github.com/rjeczalik/notify"
	"log"
	"os"
	"path/filepath"
	"time"
)

func NewWatch(syncID, path string, ignore bool) (*Watcher, error) {
	if _, err := os.Stat(path); err != nil {
		return nil, err
	}

	c := make(chan notify.EventInfo, 100)
	err := notify.Watch(filepath.Join(path, "..."), c, notify.Create, notify.Write, notify.Remove, notify.InMovedTo, notify.InMovedFrom)

	return &Watcher{
		SyncID: syncID,
		Path:   path,
		Chan:   c,
		Ignore: ignore,
	}, err
}

func (wm *WatchManger) add(w *Watcher) {

	log.Printf("ID %s 路径 %s 开始监控", w.SyncID, w.Path)
	rootPathLen := len(w.Path)
	var p string
	t := time.Now().Unix()

	PF := func(event notify.EventInfo) {

		now := time.Now().Unix()
		if p == event.Path() && now-t < 1 {
			log.Println("跳过的path", event.Path(), now, t)
			return
		}
		p = event.Path()
		t = now

		absPath := event.Path()[rootPathLen:]
		if event.Event() == notify.InMovedTo {
			for {
				select {
				case eid := <-w.Chan:
					if eid.Event() == notify.InMovedFrom {
						eidAbsPath := eid.Path()[rootPathLen:]

						wm.RenameChannel <- FsEventWithID{event.Path(), notify.Rename, w.SyncID, absPath}
						wm.RenameChannel <- FsEventWithID{eid.Path(), notify.Rename, w.SyncID, eidAbsPath}
						log.Println(event.Path(), event.Event())
						log.Println(eid.Path(), eid.Event())
						goto End
					}
				case <-time.After(time.Second):
					goto End
				}
			}
		End:
			return
		}

		if event.Event() == notify.InMovedFrom {
			wm.EventWithIDChan <- FsEventWithID{event.Path(), notify.Remove, w.SyncID, absPath}
			return
		}

		wm.EventWithIDChan <- FsEventWithID{event.Path(), event.Event(), w.SyncID, absPath}
	}

	if w.Ignore {
		for {
			event := <-w.Chan
			if _, ok := ignore.Lock.Load(event.Path()); ok {
				continue
			}

			if wm.Ignore.Match(event.Path()) {
				continue
			}

			PF(event)
		}

	}

	for { //非过滤
		event := <-w.Chan

		if _, ok := ignore.Lock.Load(event.Path()); ok {
			log.Println("切断冲突", event.Path())
			continue
		}

		PF(event)

		//now := time.Now().Unix()
		//if p == event.Path() && now-t < 1 {
		//	log.Println("跳过的path", event.Path(), now, t)
		//	continue
		//}
		//p = event.Path()
		//t = now
		//
		//log.Println(event.Event(), event.Path())
		//absPath := event.Path()[rootPathLen:]
		//
		//if event.Event() == notify.Rename {
		//	//wm.RenameChannel <- FsEventWithID{event, w.SyncID, absPath}
		//	continue
		//}
		//wm.EventWithIDChan <- FsEventWithID{event.Path(), event.Event(), w.SyncID, absPath}
		////wm.EventWithIDChan <- FsEventWithID{event, w.SyncID, absPath}
	}

}
