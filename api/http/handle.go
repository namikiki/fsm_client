package httpapi

import (
	"net/http"

	"github.com/fsnotify/fsnotify"
)

type Handle struct {
	client *http.Client
	watch  *fsnotify.Watcher
}

func New(client *http.Client, watch *fsnotify.Watcher) Handle {
	return Handle{client: client, watch: watch}
}
