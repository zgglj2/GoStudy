package main

import (
	"fmt"
	"os"
)

func main() {
	fullexecpath, err := os.Executable()
	if err != nil {
		fmt.Println("os.Executable() error: ", err)
		return
	}
	fmt.Println("os.Executable(): ", fullexecpath)
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("os.Getwd() error: ", err)
		return
	}
	fmt.Println("os.Getwd(): ", dir)
}
