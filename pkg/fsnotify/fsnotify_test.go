package fsn

import (
	"log"
	"testing"

	"fsm_client/pkg/config"
	"fsm_client/pkg/ignore"

	"github.com/fsnotify/fsnotify"
)

func TestNew(t *testing.T) {

	watchers := make(map[string]Watch)
	watcher, err := NewWatch("123", "/Users/zylzyl/go/src/fsm_client/pkg/mock/test")
	if err != nil {
		log.Println(err)
	}
	watchers["123"] = watcher

	newIgnoreConfig, err := config.NewIgnoreConfig()
	if err != nil {
		log.Println(err)
	}

	newIgnore, err := ignore.NewIgnore(newIgnoreConfig)
	if err != nil {
		log.Println(err)
	}

	buffChannal := make(chan fsnotify.Event)
	errBuffChannal := make(chan error)

	manger := NewWatchManger(watchers, newIgnore, buffChannal, errBuffChannal)
	go manger.Watch()
	manger.ProessChan()
}
