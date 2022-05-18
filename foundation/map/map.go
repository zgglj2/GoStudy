package main

import (
	"fmt"
	"sort"
)

func ModifyMap(m map[string]string) {
	m["one"] = "one"
}

func main() {
	countryCapitalMap := make(map[string]string)

	countryCapitalMap["France"] = "巴黎"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Japan"] = "东京"
	countryCapitalMap["India "] = "新德里"
	fmt.Println("原始的 map 大小是", len(countryCapitalMap))
	ModifyMap(countryCapitalMap)
	fmt.Println("修改后的 map 大小是", len(countryCapitalMap))

	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}

	capital, ok := countryCapitalMap["American"]
	if ok {
		fmt.Println("American 的首都是", capital)
	} else {
		fmt.Println("American 的首都不存在")
	}

	delete(countryCapitalMap, "France")
	fmt.Println("法国条目被删除")

	fmt.Println("删除元素后地图")

	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}

	x := map[int]int{}
	for i := 0; i < 10000; i++ {
		x[i] = i
	}
	fmt.Println("初始化后,长度:", len(x))

	// 遍历时删除所有的偶数
	for k := range x {
		if k%2 == 0 {
			delete(x, k)
		}
	}
	fmt.Println("删除所有的偶数后,长度:", len(x))

	// 遍历时删除所有的元素
	for k := range x {
		delete(x, k)
	}
	fmt.Println("删除所有的元素后,长度:", len(x))

	// Version A:
	items := make([]map[int]int, 5)
	for i := range items {
		items[i] = make(map[int]int)
		items[i][1] = 2
	}
	// Version B: NOT GOOD!
	fmt.Printf("Version A: Value of items: %v\n", items)
	items2 := make([]map[int]int, 5)
	for _, item := range items2 {
		item = make(map[int]int) // item is only a copy of the slice element.
		item[1] = 2              // This 'item' will be lost on the next iteration.
	}
	fmt.Printf("Version A: Value of items: %v\n", items2)

	barVal := map[string]int{"alpha": 34, "bravo": 56, "charlie": 23, "delta": 87, "echo": 56,
		"foxtrot": 12, "golf": 34, "hotel": 16, "indio": 87, "juliet": 65, "kili": 43, "lima": 98}
	fmt.Println("unsorted:")
	for k, v := range barVal {
		fmt.Printf("Key: %v, Value: %v / ", k, v)
	}

	keys := make([]string, 0, len(barVal))
	for k := range barVal {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	fmt.Println()
	fmt.Println("sorted:")
	for _, k := range keys {
		fmt.Printf("Key: %v, Value: %v / ", k, barVal[k])
	}
}
