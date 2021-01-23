package main

import (
	"fmt"
	"math"
)

func Add(x int, y int) int {
	return x + y
}

func Add2(x, y int) int {
	return x + y
}

func Swap(x, y string) (string, string) {
	return y, x
}

func Split(sum int) (x, y int) {
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

func min(s ... int) int {
	if len(s) == 0 {
		return 0
	}
	min := s[0]
	for _, v := range s {
		if v < min {
			min = v
		}
	}
	return min
}

func f() (ret int) {
	defer func() {
		ret++
	}()
	return 1
}

func Add3() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}

func Adder3(a int) func(b int) int {
	return func(b int) int {
		return a + b
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

	x := min(1, 3, 2, 0)
	fmt.Printf("The minimum is: %d\n", x)
	slice := []int{7, 9, 3, 5, 1}
	x = min(slice...)
	fmt.Printf("The minimum in the slice is: %d\n", x)

	fmt.Println(f())

	p3 := Add3()
	fmt.Printf("Call Add3 for 3 gives: %v\n", p3(3))
	TwoAdder := Adder3(2)
	fmt.Printf("The result is: %v\n", TwoAdder(3))
}
