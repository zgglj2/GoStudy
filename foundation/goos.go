package main

import (
	"runtime"
	"fmt"
	"os"
)

func main() {
	osinfo := runtime.GOOS
	fmt.Println(osinfo)

	path := os.Getenv("PATH")
	fmt.Println(path)
}
