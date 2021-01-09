package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("os: ", runtime.GOOS)
	fmt.Println("os: ", runtime.GOARCH)
	fmt.Println("os: ", runtime.GOROOT())
}
