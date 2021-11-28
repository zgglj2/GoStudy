package taillog

import (
	"GoStudy/logagent/etcd"
	"context"
	"fmt"
)

var (
	tailtaskMgr *TailTaskMgr
)

type TailTaskMgr struct {
	tailTaskMap map[string]*TailTask
	confChan    <-chan []*etcd.LogEntry
	ctx         context.Context
	cancel      context.CancelFunc
}

func InitTailTaskMgr(logEntryConf []*etcd.LogEntry, confChan <-chan []*etcd.LogEntry) {
	tailtaskMgr = &TailTaskMgr{
		tailTaskMap: make(map[string]*TailTask, 32),
		confChan:    confChan,
	}

	for index, value := range logEntryConf {
		fmt.Printf("index: %v, value: %v\n", index, value)
		tailTask, err := NewTailTask(value.Path, value.Topic)
		if err != nil {
			fmt.Printf("new tailtask failed, err: %v\n", err)
			continue
		}
		tailtaskMgr.tailTaskMap[value.Path] = tailTask

	}
	tailtaskMgr.ctx, tailtaskMgr.cancel = context.WithCancel(context.Background())
	go run()
}

func Finish() {
	tailtaskMgr.cancel()
	for _, tailtask := range tailtaskMgr.tailTaskMap {
		tailtask.Finish()
	}

}
func exist(path string) bool {
	for taskPath := range tailtaskMgr.tailTaskMap {
		if taskPath == path {
			return true
		}
	}
	return false
}

func run() {
	for {
		select {
		case <-tailtaskMgr.ctx.Done():
			fmt.Printf("taillog manager is stop...\n")
			return
		case newConf := <-tailtaskMgr.confChan:
			fmt.Printf("%#v\n", newConf)
			for _, conf := range newConf {
				if exist(conf.Path) {
					fmt.Printf("the task of path:%s is already run\n", conf.Path)
					continue
				}
				tailTask, err := NewTailTask(conf.Path, conf.Topic)
				if err != nil {
					fmt.Printf("new tailtask failed, err: %v\n", err)
					continue
				}
				fmt.Printf("create new task for path %v, topic: %v success\n", conf.Path, conf.Topic)
				tailtaskMgr.tailTaskMap[conf.Path] = tailTask

			}
			for path, tailTask := range tailtaskMgr.tailTaskMap {
				isFound := false
				for _, conf := range newConf {
					if path == conf.Path {
						isFound = true
						break
					}
				}
				if !isFound {
					tailTask.cancel()
					delete(tailtaskMgr.tailTaskMap, path)
					fmt.Printf("the task of path:%s is remove from tailTaskMap\n", path)
				}
			}
		}
	}
}
