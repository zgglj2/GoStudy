package main

import (
	"GoStudy/foundation/packutil"
	"fmt"
)

func main() {
	test1 := packutil.ReturnStr()
	fmt.Printf("ReturnStr from package1: %s\n", test1)
	fmt.Printf("Integer from package1: %d\n", packutil.Pack1Int)
	// fmt.Printf("Float from package1: %f\n", pack1.pack1Float)
}
