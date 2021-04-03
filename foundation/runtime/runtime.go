package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("os: ", runtime.GOOS)
	fmt.Println("os: ", runtime.GOARCH)
	fmt.Println("os: ", runtime.GOROOT())

	where := func() {
		pc, file, line, ok := runtime.Caller(1)
		fmt.Printf("%s:%d\n", file, line)
		fmt.Printf("%v:%v", pc, ok)
	}
	where()
}
