package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

func logicCode() {
	var c chan int
	for {
		select {
		case v := <-c:
			fmt.Printf("recv from chan, value:%v\n", v)
		default:
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func main() {
	var isCPUPprof bool
	var isMemPprof bool

	flag.BoolVar(&isCPUPprof, "cpu", false, "turn cpu pprof on")
	flag.BoolVar(&isMemPprof, "mem", false, "turn mem pprof on")
	flag.Parse()

	if isCPUPprof {
		file, err := os.Create("./cpu.prof")
		if err != nil {
			fmt.Println("create cpu pprof failed, err: ", err)
			return
		}
		defer file.Close()
		pprof.StartCPUProfile(file)
		defer pprof.StopCPUProfile()
	}
	for i := 0; i < 6; i++ {
		go logicCode()
	}
	time.Sleep(time.Second * 20)
	if isMemPprof {
		file, err := os.Create("./mem.prof")
		if err != nil {
			fmt.Println("create mem pprof failed, err: ", err)
			return
		}
		pprof.WriteHeapProfile(file)
		file.Close()
	}

}
