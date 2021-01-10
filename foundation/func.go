package main

import (
	"fmt"
	"math"
)

func Add(x int, y int) int {
	return x + y
}

func Add2(x , y int) int {
	return x + y
}

func Swap(x, y string) (string, string) {
	return y, x
}

func Split(sum int)(x, y int) {
	x = sum * 4 / 9;
	y = sum - x
	return
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	fmt.Println("x + y = ", Add(1, 2))
	fmt.Println("x + y = ", Add2(3, 4))
	a, b := Swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(Split(17))

	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
