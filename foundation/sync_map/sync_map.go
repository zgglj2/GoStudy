package main

import (
	"fmt"
	"strconv"
	"sync"
)

// var m1 = make(map[string]int)

// func get(key string) int {
// 	return m1[key]
// }

// func set(key string, value int) {
// 	m1[key] = value
// }

// func not_sync_map() {
// 	wg := sync.WaitGroup{}
// 	for i := 0; i < 20; i++ {
// 		wg.Add(1)
// 		go func(n int) {
// 			key := strconv.Itoa(n)
// 			set(key, n)
// 			fmt.Printf("k=:%v,v:=%v\n", key, get(key))
// 			wg.Done()
// 		}(i)
// 	}
// 	wg.Wait()
// }

var m = sync.Map{}

func sync_map() {
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m.Store(key, n)
			value, _ := m.Load(key)
			fmt.Printf("k=:%v,v:=%v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func main() {
	// not_sync_map()
	sync_map()
}
