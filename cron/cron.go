package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/robfig/cron"
)

func main() {
	c := cron.New()
	c.AddFunc("@every 1s", func() { fmt.Println("tick every 1 second run once by @every 1s") })
	c.AddFunc("0/1 * * * * *", func() { fmt.Println("tick every 1 second run once by 0/1") })
	c.AddFunc("0 30 * * * *", func() { fmt.Println("Every hour on the half hour") })
	c.AddFunc("@hourly", func() { fmt.Println("Every hour") })
	c.AddFunc("@every 1h30m", func() { fmt.Println("Every hour thirty") })
	c.Start()
	defer c.Stop()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT)
	<-ch
}
