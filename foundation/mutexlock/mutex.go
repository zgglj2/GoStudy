package main

import (
	"fmt"
	"sync"
	"time"
)

// 共享变量
var (
	m  sync.Mutex
	v1 int
)

// 修改共享变量
// 在Lock()和Unlock()之间的代码部分是临界区
func change(i int) {
	m.Lock()
	time.Sleep(time.Second)
	v1 = v1 + 1
	if v1%10 == 0 {
		v1 = v1 - 10*i
	}
	m.Unlock()
}

// 访问共享变量
// 在Lock()和Unlock()之间的代码部分是是临界区
func read() int {
	m.Lock()
	a := v1
	m.Unlock()
	return a
}

func main() {
	var numGR = 21
	var wg sync.WaitGroup

	fmt.Printf("%d", read())

	// 循环创建numGR个goroutine
	// 每个goroutine都执行change()、read()
	// 每个change()和read()都会持有锁
	for i := 0; i < numGR; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			change(i)
			fmt.Printf(" -> %d", read())
		}(i)
	}

	wg.Wait()
}
