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
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d", file, line)
	}
	where()
}
