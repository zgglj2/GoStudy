package taillog

import (
	"GoStudy/logagent/kafka"
	"context"
	"fmt"

	"github.com/nxadm/tail"
)

type TailTask struct {
	path     string
	topic    string
	instance *tail.Tail
	ctx      context.Context
	cancel   context.CancelFunc
}

func NewTailTask(path, topic string) (tailTask *TailTask, err error) {
	tailTask = &TailTask{
		path:  path,
		topic: topic,
	}
	err = tailTask.Init()
	return
}

func (t *TailTask) Init() (err error) {
	config := tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	}
	t.instance, err = tail.TailFile(t.path, config)
	if err != nil {
		fmt.Println("tail file fail, err: ", err)
		return
	}
	t.ctx, t.cancel = context.WithCancel(context.Background())
	go t.run()
	return
}

func (t *TailTask) Finish() {
	t.cancel()
	t.instance.Cleanup()
}

func (t *TailTask) run() {
	for {
		select {
		case <-t.ctx.Done():
			fmt.Printf("the task for path:%s is stop...\n", t.path)
			return
		case line := <-t.instance.Lines:
			kafka.SendToChan(t.topic, line.Text)
		}
	}
}
