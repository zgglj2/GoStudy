package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type singleton struct{}

var instance *singleton

var initialized uint32

var mu sync.Mutex
var once sync.Once

func GetInstance1() *singleton {
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}

	mu.Lock()
	defer mu.Unlock()

	if initialized == 0 {
		instance = &singleton{}
		atomic.StoreUint32(&initialized, 1)
	}
	return instance
}
func GetInstance2() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}

func main() {
	fmt.Printf("%#v\n", GetInstance1())
	fmt.Printf("%#v\n", GetInstance2())
}
