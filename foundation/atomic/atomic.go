package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	count int32
	wg    sync.WaitGroup
)

func incCount() {
	defer wg.Done()

	for i := 0; i < 2; i++ {
		value := atomic.LoadInt32(&count)
		runtime.Gosched()
		value++
		atomic.StoreInt32(&count, value)
	}
}

func main() {
	wg.Add(2)

	go incCount()
	go incCount()

	wg.Wait()
	fmt.Println(count)

}
