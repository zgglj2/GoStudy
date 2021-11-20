package main

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

func main() {
	// 设置日志输出为os.Stdout
	logger.Out = os.Stdout

	//可以设置像文件等任意`io.Writer`类型作为日志输出
	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logger.Info("Failed to log to file, using default stderr")
	}
	logger.Out = io.MultiWriter(os.Stdout, file)
	logger.WithFields(logrus.Fields{
		"animal": "dog",
	}).Info("一条舔狗出现了。")

	logger.SetLevel(logrus.TraceLevel)
	logger.SetReportCaller(true)
	logger.Trace("Something very low level.")
	logger.Debug("Useful debugging information.")
	logger.Info("Something noteworthy happened!")
	logger.Warn("You should probably take a look at this.")
	logger.Error("Something failed but I'm not quitting.")
	// 记完日志后会调用os.Exit(1)
	// logger.Fatal("Bye.")
	// 记完日志后会调用 panic()
	// logger.Panic("I'm bailing.")

	requestLogger := logger.WithFields(logrus.Fields{"name": "glj", "user_ip": "1.1.1.1"})
	requestLogger.Info("something happened on that request")
	requestLogger.Warn("something not great happened")
}
