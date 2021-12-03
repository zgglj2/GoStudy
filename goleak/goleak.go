package main

import (
	"fmt"
	"time"
)

func leak() error {
	go func() {
		time.Sleep(time.Minute)
	}()

	return nil
}

func leak2() error {
	ch := make(chan int)

	go func() {
		val := <-ch
		fmt.Println("We received a value:", val)
	}()
	return nil
}

func main() {
	if err := leak(); err != nil {
		fmt.Println("leak() error not expected")
	}
	if err := leak2(); err != nil {
		fmt.Println("leak2() error not expected")
	}
}
