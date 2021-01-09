package main

import "fmt"

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

func main() {
	fmt.Println("x + y = ", Add(1, 2))
	fmt.Println("x + y = ", Add2(3, 4))
	a, b := Swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(Split(17))
}
