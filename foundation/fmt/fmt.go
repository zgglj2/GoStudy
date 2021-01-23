package main

import "fmt"

func main() {
	n := 100
	fmt.Printf("%T\n", n)
	fmt.Printf("%v\n", n)
	fmt.Printf("%b\n", n)
	fmt.Printf("%d\n", n)
	fmt.Printf("%o\n", n)
	fmt.Printf("%x\n", n)

	s := "Hello"
	fmt.Printf("string: %s\n", s)
	fmt.Printf("string: %v\n", s)
	fmt.Printf("string: %#v\n", s)
}
