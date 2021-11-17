package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"syscall"
	"time"
)

// func TryE() {
// 	errs := recover()
// 	if errs == nil {
// 		return
// 	}
// 	exeName := os.Args[0]                                             //获取程序名称
// 	now := time.Now()                                                 //获取当前时间
// 	pid := os.Getpid()                                                //获取进程ID
// 	time_str := now.Format("20060102150405")                          //设定时间格式
// 	fname := fmt.Sprintf("%s-%d-%s-dump.log", exeName, pid, time_str) //保存错误信息文件名:程序名-进程ID-当前时间（年月日时分秒）
// 	fmt.Println("dump to file ", fname)
// 	f, err := os.Create(fname)
// 	if err != nil {
// 		return
// 	}
// 	defer f.Close()
// 	f.WriteString(fmt.Sprintf("%v\r\n", errs)) //输出panic信息
// 	f.WriteString("========\r\n")
// 	f.WriteString(string(debug.Stack())) //输出堆栈信息
// }

var (
	kernel32         = syscall.MustLoadDLL("kernel32.dll")
	procSetStdHandle = kernel32.MustFindProc("SetStdHandle")
)

func setStdHandle(stdhandle int32, handle syscall.Handle) error {
	r0, _, e1 := syscall.Syscall(procSetStdHandle.Addr(), 2, uintptr(stdhandle), uintptr(handle), 0)
	if r0 == 0 {
		if e1 != 0 {
			return error(e1)
		}
		return syscall.EINVAL
	}
	return nil
}

// redirectStderr to the file passed in
func redirectStderr(f *os.File) {
	if runtime.GOOS == "windows" {
		err := setStdHandle(syscall.STD_ERROR_HANDLE, syscall.Handle(f.Fd()))
		if err != nil {
			log.Fatalf("Failed to redirect stderr to file: %v", err)
		}
		// SetStdHandle does not affect prior references to stderr
		os.Stderr = f
	} else {
		// err := syscall.Dup2(int(f.Fd()), int(os.Stderr.Fd()))
		// if err != nil {
		// 	log.Fatalf("Failed to redirect stderr to file: %v", err)
		// }
	}

}

func saferoutine(c chan bool) {
	for i := 0; i < 10; i++ {
		fmt.Println("Count:", i)
		time.Sleep(1 * time.Second)
	}
	c <- true
}
func panicgoroutine(c chan bool) {
	time.Sleep(5 * time.Second)

	panic("Panic, omg ...")
	// c <- true
}
func main() {
	logFile, err := os.OpenFile("./fatal.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0660)
	if err != nil {
		log.Println("服务启动出错", "打开异常日志文件失败", err)
		return
	}
	redirectStderr(logFile)

	c := make(chan bool, 2)
	go saferoutine(c)
	go panicgoroutine(c)
	for i := 0; i < 2; i++ {
		<-c
	}
}
