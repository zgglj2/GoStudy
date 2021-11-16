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

func DumpKeyValue(key, value interface{}) bool {
	fmt.Printf("  k=:%v,v:=%v\n", key, value)
	return true
}
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

	value2, _ := m.LoadOrStore("test", 100)
	fmt.Printf("k=:%v,v:=%v\n", "test", value2)

	value3, _ := m.LoadAndDelete("1")
	fmt.Printf("k=:%v,v:=%v\n", "1", value3)

	m.Delete("0")

	m.Range(DumpKeyValue)

	fmt.Println(m)
}

func main() {
	// not_sync_map()
	sync_map()
}
