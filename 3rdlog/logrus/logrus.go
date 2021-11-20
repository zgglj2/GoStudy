package main

import (
	"io"
	"os"

	log "github.com/sirupsen/logrus"
)

type DefaultFieldHook struct {
}

func (hook *DefaultFieldHook) Fire(entry *log.Entry) error {
	// fmt.Println(entry.Level.String())
	if entry.Level.String() == "warning" {
		entry.Message = "hook message"
	}
	if entry.Data["name"] == "glj" {
		entry.Data["name"] = "jrj"
	}
	return nil
}

func (hook *DefaultFieldHook) Levels() []log.Level {
	return log.AllLevels
}

var logger = log.New()

func main() {

	// 设置日志输出为os.Stdout
	logger.Out = os.Stdout

	//可以设置像文件等任意`io.Writer`类型作为日志输出
	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logger.Info("Failed to log to file, using default stderr")
	}
	logger.Out = io.MultiWriter(os.Stdout, file)
	logger.WithFields(log.Fields{
		"animal": "dog",
	}).Info("一条舔狗出现了。")

	logger.SetLevel(log.TraceLevel)
	logger.SetReportCaller(true)

	var hook DefaultFieldHook
	logger.AddHook(&hook)

	logger.Trace("Something very low level.")
	logger.Debug("Useful debugging information.")
	logger.Info("Something noteworthy happened!")
	logger.Warn("You should probably take a look at this.")
	logger.Error("Something failed but I'm not quitting.")
	// 记完日志后会调用os.Exit(1)
	// logger.Fatal("Bye.")
	// 记完日志后会调用 panic()
	// logger.Panic("I'm bailing.")

	requestLogger := logger.WithFields(log.Fields{"name": "glj", "user_ip": "1.1.1.1"})
	requestLogger.Info("something happened on that request")
	requestLogger.Warn("something not great happened")
}
