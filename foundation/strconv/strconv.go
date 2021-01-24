package main

import (
	"fmt"
	"strconv"
)

func main() {
	n, err := strconv.ParseInt("1234", 10, strconv.IntSize)
	fmt.Println(n, err)

	v32 := "-354634382"
	if s, err := strconv.ParseInt(v32, 10, 32); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}
	v64 := "-3546343826724305832"
	if s, err := strconv.ParseInt(v64, 10, 64); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}

	v := "10"
	if s, err := strconv.Atoi(v); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}

	n = 4321024
	ss := strconv.FormatInt(n, 10)
	fmt.Printf("%T, %v\n", ss, ss)

	nn := strconv.Itoa(1024)
	fmt.Printf("%T, %v\n", nn, nn)
}
