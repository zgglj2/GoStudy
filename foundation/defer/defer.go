package main

import (
	"fmt"
	"io"
)

func func1(s string) (n int, err error) {
	defer func() {
		fmt.Printf("func1(%q) = %d, %v\n", s, n, err)
	}()
	return 7, io.EOF
}

func main() {
	defer fmt.Println("world")
	fmt.Println("hello")

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")

	func1("Golang")
}
