package main

import "fmt"

func main() {
	var numbers1 = make([]int,3,5)

	printSlice(numbers1)

	var numbers2 []int

	printSlice(numbers2)

	numbers := []int{0,1,2,3,4,5,6,7,8}
	printSlice(numbers)

	fmt.Println("numbers ==", numbers)

	fmt.Println("numbers[1:4] ==", numbers[1:4])

	fmt.Println("numbers[:3] ==", numbers[:3])

	fmt.Println("numbers[4:] ==", numbers[4:])

	numbers4 := make([]int,0,5)
	printSlice(numbers4)

	number2 := numbers[:2]
	printSlice(number2)

	number3 := numbers[2:5]
	printSlice(number3)


	var numbers5 []int
	printSlice(numbers5)

	numbers5 = append(numbers5, 0)
	printSlice(numbers5)

	numbers5 = append(numbers5, 1)
	printSlice(numbers5)

	numbers5 = append(numbers5, 2,3,4)
	printSlice(numbers5)

	numbers6 := make([]int, len(numbers5), (cap(numbers5))*2)

	copy(numbers6,numbers5)
	printSlice(numbers6)

}

func printSlice(x []int){
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}