package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/stevedonovan/luar"
	"gopkg.in/fsnotify.v1"
)

var (
	watcher *fsnotify.Watcher
	err     error

	ctx    context.Context
	cancel context.CancelFunc
)

func SetWatchPath(path []string) {
	for _, v := range path {
		fmt.Println(v)
		err = watcher.Add(v)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func StartFileMonitor() {
	watcher, err = fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel = context.WithCancel(context.Background())
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Printf("file monitor is stop...\n")
				return
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
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

}

func StopFileMonitor() {
	cancel()
}

func MySleep(sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
}
func main() {
	L := luar.Init()
	defer L.Close()

	luar.Register(L, "", luar.Map{
		"print":            fmt.Println,
		"sleep":            time.Sleep,
		"StartFileMonitor": StartFileMonitor,
		"StopFileMonitor":  StopFileMonitor,
		"SetWatchPath":     SetWatchPath,
	})
	// executes a file
	if err := L.DoFile("fsnotify.lua"); err != nil {
		panic(err)
	}

	<-ctx.Done()
}
