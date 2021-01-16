package main;

import "C"

func main() {}

//export Hello
func Hello() string {
	return "Hello"
}

//export Print_a
func Print_a() {
	println("aaaaa")
}