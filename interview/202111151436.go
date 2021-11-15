package main

import "fmt"

func GetValue(m map[int]string, id int) (string, bool) {
	if _, exist := m[id]; exist {
		return "存在数据", true
	}
	return "", false
}

func main() {
	fmt.Println(len("你好cd!"))

	intmap := map[int]string{
		1: "a",
		2: "bb",
		3: "ccc",
	}
	v, ok := GetValue(intmap, 3)
	fmt.Println(len(intmap), v, ok)
}
