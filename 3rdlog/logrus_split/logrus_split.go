package main

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var path = "logrus_split.log"

var logger = log.New()

func init() {
	writer := &lumberjack.Logger{
		Filename:   path,
		MaxSize:    1, // megabytes
		MaxBackups: 3,
		// MaxAge:     28,   //days
		Compress: true, // disabled by default
	}

	// logger.Out = io.MultiWriter(os.Stdout, writer)
	logger.Out = writer
	//log.SetFormatter(&log.JSONFormatter{})

	logger.WithFields(log.Fields{
		"animal": "dog",
	}).Info("一条舔狗出现了。")
}
func main() {
	for {
		logger.Info("hello, world!")
		// time.Sleep(time.Duration(2) * time.Millisecond * 1)
	}
}
