package main

import (
	"flag"
	"log"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
)

var hello []int

func counter(wg *sync.WaitGroup) {
	defer wg.Done()

	slice := []int{0}
	c := 1
	for i := 0; i < 100000; i++ {
		c = i + 1 + 2 + 3 + 4 + 5
		slice = append(slice, c)
	}
	hello = slice
}

func main() {
	runtime.GOMAXPROCS(1)

	var traceProfile = flag.String("traceprofile", "trace.pprof", "write trace profile to file")
	flag.Parse()

	if *traceProfile != "" {
		f, err := os.Create(*traceProfile)
		if err != nil {
			log.Fatal(err)
		}
		trace.Start(f)
		defer f.Close()
		defer trace.Stop()
	}

	var wg sync.WaitGroup
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go counter(&wg)
	}
	wg.Wait()
}
