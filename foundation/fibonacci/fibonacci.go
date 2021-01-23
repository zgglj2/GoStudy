package main

import (
	"time"
	"fmt"
)

const LIM = 41

var fibs [LIM]uint64

func fibonacci3(n int) (res uint64) {
	// memoization: check if fibonacci(n) is already known in array:
	if fibs[n] != 0 {
		res = fibs[n]
		return
	}
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci3(n-1) + fibonacci3(n-2)
	}
	fibs[n] = res
	return
}

func main() {
	var result uint64 = 0
	start := time.Now()
	for i := 0; i < LIM; i++ {
		result = fibonacci3(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
}
