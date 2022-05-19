package main

import (
	"fmt"
	"strconv"
)

func AnyTest(name string, data any) error {
	fmt.Printf("data.(type) = %v", data)
	switch data.(type) {
	case string:
		fmt.Println(name, data.(string))
	case int:
		fmt.Println(name, strconv.Itoa(data.(int)))
	}

	return nil
}

func main() {
	AnyTest("zhang", "1")
	AnyTest("zhang", 2)
}
