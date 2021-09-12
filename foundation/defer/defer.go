package main

import (
	"fmt"
	"io"
	"log"
	"time"
)

func func1(s string) (n int, err error) {
	defer func() {
		fmt.Printf("func1(%q) = %d, %v\n", s, n, err)
	}()
	return 7, io.EOF
}
func bigSlowOperation() {
	defer trace("bigSlowOperation")() // don't forget the extra parentheses
	// ...lots of work…
	time.Sleep(10 * time.Second) // simulate slow operation by sleeping
}
func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
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

	bigSlowOperation()
}
