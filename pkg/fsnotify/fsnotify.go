package fsn

import (
	"log"

	"github.com/fsnotify/fsnotify"
)

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
