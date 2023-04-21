//go:build darwin

package fsn

func NewWatch(syncID, path string, ignore bool) (*Watcher, error) {
	if _, err := os.Stat(path); err != nil {
		return nil, err
	}

	c := make(chan notify.EventInfo, 100)
	err := notify.Watch(filepath.Join(path, "..."), c, notify.All)

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
	PF := func() {

		now := time.Now().Unix()
		if p == event.Path() && now-t < 1 {
			log.Println("跳过的path", event.Path(), now, t)
			return
		}
		p = event.Path()
		t = now

		absPath := event.Path()[rootPathLen:]

		if event.Event() == notify.Rename {
			wm.RenameChannel <- FsEventWithID{event, w.SyncID, absPath}
			return
		}

		wm.EventWithIDChan <- FsEventWithID{event, w.SyncID, absPath}

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

		log.Printf("ID %s 路径 %s 停止监控", w.SyncID, w.Path)
		return
	}

	for { //非过滤
		event := <-w.Chan

		if _, ok := ignore.Lock.Load(event.Path()); ok {
			log.Println("切断冲突", event.Path())
			continue
		}

		PF(event)
	}

}
