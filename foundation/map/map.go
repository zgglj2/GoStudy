package main

import "fmt"

func main() {
	countryCapitalMap := make(map[string]string)

	countryCapitalMap["France"] = "巴黎"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Japan"] = "东京"
	countryCapitalMap["India "] = "新德里"

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
}
