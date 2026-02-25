package main

import (
	"strings"
	"sync"

	"github.com/fsnotify/fsnotify"
)

type FileWatcher struct {
	watcher  *fsnotify.Watcher
	notesDir string
	mu       sync.RWMutex
	callback func(string)
}

var fileWatcher *FileWatcher

func StartFileWatcher(dir string, callback func(string)) error {
	if fileWatcher != nil {
		fileWatcher.Stop()
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}

	fileWatcher = &FileWatcher{
		watcher:  watcher,
		notesDir: dir,
		callback: callback,
	}

	err = watcher.Add(dir)
	if err != nil {
		return err
	}

	go fileWatcher.watch()

	return nil
}

func (fw *FileWatcher) watch() {
	for {
		select {
		case event, ok := <-fw.watcher.Events:
			if !ok {
				return
			}

			if event.Op&fsnotify.Create == fsnotify.Create ||
				event.Op&fsnotify.Write == fsnotify.Write ||
				event.Op&fsnotify.Remove == fsnotify.Remove {

				if strings.HasSuffix(event.Name, ".md") && fw.callback != nil {
					fw.callback(event.Name)
				}
			}
		case err, ok := <-fw.watcher.Errors:
			if !ok {
				return
			}
			_ = err
		}
	}
}

func (fw *FileWatcher) Stop() {
	if fw != nil && fw.watcher != nil {
		fw.watcher.Close()
	}
}