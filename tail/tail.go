package main

import (
	"fmt"
	"time"

	"github.com/hpcloud/tail"
)

func main() {
	config := tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	}
	tails, err := tail.TailFile("my.log", config)
	if err != nil {
		fmt.Println("tail file fail, err: ", err)
		return
	}
	var line *tail.Line
	var ok bool
	for {
		line, ok = <-tails.Lines
		if !ok {
			fmt.Println("tail file close reopen, filename: ", tails.Filename)
			time.Sleep(time.Second)
			continue
		}
		fmt.Println("line: ", line.Text)
	}
}
