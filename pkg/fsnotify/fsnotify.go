package fsn

import (
	"log"

	"fsm_client/pkg/ignore"

	"github.com/rjeczalik/notify"
)

type Watcher struct {
	Chan   chan notify.EventInfo
	SyncID string
	Path   string
	Ignore bool
}

type FsEventWithID struct {
	Path string
	notify.Event
	SyncID  string
	AbsPath string
}

type WatchManger struct {
	Watchers            map[string]*Watcher
	EventWithIDChan     chan FsEventWithID
	ErrBuffChannel      chan error
	RenameChannel       chan FsEventWithID
	AddNotifyChannel    chan *Watcher
	RemoveNotifyChannel chan string
	Ignore              *ignore.Ignore
}

// NewWatchManger  NewWatchManger
func NewWatchManger(buffLen int64, ignore *ignore.Ignore,
) *WatchManger {

	eventWithIDChan := make(chan FsEventWithID, buffLen)
	RenameChannel := make(chan FsEventWithID, buffLen)
	errBuffChannel := make(chan error, 2)
	removeChannel := make(chan string, 2)
	addChannel := make(chan *Watcher, 4)
	watchers := make(map[string]*Watcher)

	return &WatchManger{
		Watchers:            watchers,
		AddNotifyChannel:    addChannel,
		RenameChannel:       RenameChannel,
		RemoveNotifyChannel: removeChannel,
		EventWithIDChan:     eventWithIDChan,
		ErrBuffChannel:      errBuffChannel,
		Ignore:              ignore,
	}
}

func (wm *WatchManger) remove(syncID string) {
	notify.Stop(wm.Watchers[syncID].Chan)
	delete(wm.Watchers, syncID)
}

func (wm *WatchManger) Watch() {

	for {
		select {
		case a := <-wm.AddNotifyChannel:
			if _, ok := wm.Watchers[a.SyncID]; ok {
				log.Println("已监测", a.Path)
				continue
			}

			wm.Watchers[a.SyncID] = a
			go wm.add(a)
		case r := <-wm.RemoveNotifyChannel:
			wm.remove(r)
		}
	}

}
