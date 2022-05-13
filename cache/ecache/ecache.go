package main

import (
	"fmt"
	"time"

	"github.com/orca-zhang/ecache"
)

type UserInfo struct {
	Name string
	Age  int
}

func main() {
	// 创建缓存
	cache := ecache.NewLRUCache(16, 200, 3*time.Second)

	// 设置缓存
	cache.Put("key1", "value1")
	cache.Put("key2", "value2")
	cache.Put("key3", "value3")

	// 删除缓存
	cache.Del("key2")
	// 获取缓存
	if v, ok := cache.Get("key1"); ok {
		fmt.Printf("%v\n", v) //value1
	} else {
		fmt.Println("key1 not found")
	}
	if v, ok := cache.Get("key2"); ok {
		fmt.Printf("%v\n", v)
	} else {
		fmt.Println("key2 not found") //key2 not found
	}
	if v, ok := cache.Get("key3"); ok {
		fmt.Printf("%v\n", v) //value3
	} else {
		fmt.Println("key3 not found")
	}

	user := UserInfo{
		Name: "zhangsan",
		Age:  18,
	}
	cache.Put("user1", user)

	// 获取缓存
	if v, ok := cache.Get("user1"); ok {
		fmt.Printf("%v\n", v) // {zhangsan 18}
	} else {
		fmt.Println("user1 not found")
	}

	time.Sleep(5 * time.Second)
	if v, ok := cache.Get("user1"); ok {
		fmt.Printf("%v\n", v)
	} else {
		fmt.Println("user1 not found") // not found
	}

}
