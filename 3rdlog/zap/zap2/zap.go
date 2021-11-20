package main

import (
	"io"
	"os"
	"time"

	// "github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger *zap.SugaredLogger

//日志文件切割
func getWriter(filename string) io.Writer {
	return &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    1, // megabytes
		MaxBackups: 3,
		// MaxAge:     28,   //days
		Compress: true, // disabled by default
	}
}

//查看文件/文件夹是否存在
func Exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

//初始化Encoder
func initEncoder() zapcore.Encoder {
	return zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		MessageKey:  "msg",
		LevelKey:    "level",
		TimeKey:     "time",
		CallerKey:   "file",
		EncodeLevel: zapcore.CapitalLevelEncoder, //基本zapcore.LowercaseLevelEncoder。将日志级别字符串转化为小写
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		EncodeCaller: zapcore.ShortCallerEncoder, //一般zapcore.ShortCallerEncoder，以包/文件:行号 格式化调用堆栈
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) { //一般zapcore.SecondsDurationEncoder,执行消耗的时间转化成浮点型的秒
			enc.AppendInt64(int64(d) / 1000000)
		},
	})
}

func initLogger() {
	encoder := initEncoder()

	// 想要将日常文件区分开来，可以实现多个日志等级接口
	/*infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.WarnLevel
	})*/
	debugLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.DebugLevel
	})

	// 获取 info、warn日志文件的io.Writer
	//infoIoWriter := getWriter("D:/bian/logs/dspcollect.log")
	// warnIoWriter := io.MultiWriter(os.Stdout, getWriter("./zap.log"))
	warnIoWriter := getWriter("./zap.log")

	// 创建Logger
	core := zapcore.NewTee(
		//zapcore.NewCore(encoder, zapcore.AddSync(infoIoWriter), infoLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(warnIoWriter), debugLevel),
	)
	logger := zap.New(core, zap.AddCaller()) // 需要传入 zap.AddCaller() 才会显示打日志点的文件名和行数
	Logger = logger.Sugar()
}

func init() {
	initLogger()
}

func main() {
	Logger.Error("error....")
	Logger.Info("info....")
	Logger.Warn("warn....")
	for {
		Logger.Info("hello, world!")
		// time.Sleep(time.Duration(2) * time.Millisecond * 1)
	}
}
