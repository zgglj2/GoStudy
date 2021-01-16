package main

import "fmt"

type T struct {
	a int
	b float64
	c string
}

func main() {
	t := &T{7, 2.5, "abc"}
	fmt.Printf("%v\n", t)
	fmt.Printf("%+v\n", t)
	fmt.Printf("%#v\n", t)
	fmt.Printf("%T\n", t)
}
