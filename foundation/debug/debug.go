package main

import (
	"fmt"
	"runtime/debug"
	"sync"
	"time"
)

func main() {
	debug.SetMaxThreads(2)
	wg := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println(i)

			time.Sleep(time.Second * 3)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
