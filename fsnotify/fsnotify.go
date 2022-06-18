package main

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/fsnotify.v1"
)

type Event fsnotify.Event

type RecursiveWatcher struct {
	*fsnotify.Watcher
	Files   chan string
	Folders chan string
}

func (e Event) String() string {
	return fsnotify.Event(e).String()
}

func NewRecursiveWatcher(path string) (*RecursiveWatcher, error) {
	folders := Subfolders(path)
	if len(folders) == 0 {
		return nil, errors.New("No folders to watch.")
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}
	rw := &RecursiveWatcher{Watcher: watcher}

	rw.Files = make(chan string, 10)
	rw.Folders = make(chan string, len(folders))

	for _, folder := range folders {
		if err = rw.AddFolder(folder); err != nil {
			return nil, err
		}
	}
	return rw, nil
}

func (watcher *RecursiveWatcher) AddFolder(folder string) error {
	if err := watcher.Add(folder); err != nil {
		return err
	}
	watcher.Folders <- folder
	return nil
}

// Subfolders returns a slice of subfolders (recursive), including the folder provided.
func Subfolders(path string) (paths []string) {
	filepath.Walk(path, func(newPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			name := info.Name()
			// skip folders that begin with a dot
			if ShouldIgnoreFile(name) && name != "." && name != ".." {
				return filepath.SkipDir
			}
			paths = append(paths, newPath)
		}
		return nil
	})
	return paths
}

// ShouldIgnoreFile determines if a file should be ignored.
// File names that begin with "." or "_" are ignored by the go tool.
func ShouldIgnoreFile(name string) bool {
	return strings.HasPrefix(name, ".") || strings.HasPrefix(name, "_")
}

func filePathFromEvent(event *fsnotify.Event) (path string, err error) {
	defer func() { path = filepath.Clean(path) }()

	if filepath.IsAbs(event.Name) {
		path = event.Name
		return
	}

	path, err = filepath.Abs(event.Name)
	return
}

func main() {

	// watcher, err := fsnotify.NewWatcher()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer watcher.Close()

	// err = watcher.Add("/etc")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	watcher, err := NewRecursiveWatcher("/etc")
	if err != nil {
		log.Fatalf("Could not initialize fsnotify watcher: %v\n", err)
	}

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)

				if strings.TrimSpace(event.Name) == "" {
					log.Println("Got an empty string... not touching that")
					continue
				}

				path, err := filePathFromEvent(&event)
				if err != nil {
					log.Println(err)
					continue
				}
				log.Println("path:", path)

				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
				} else if event.Op&fsnotify.Remove == fsnotify.Remove {
					log.Println("Remove event received:", event.Name)
				} else if event.Op&fsnotify.Create == fsnotify.Create {
					log.Println("Create event received:", event.Name)
					stat, err := os.Stat(path)
					if err != nil {
						log.Println(err)
						continue
					}

					if stat.IsDir() {
						if err := watcher.Add(path); err != nil {
							log.Println("Couldn't watch folder:", err)
						}
					}
				} else if event.Op&fsnotify.Rename == fsnotify.Rename {
					log.Println("Rename event received:", event.Name)
				}

			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	<-done
}
