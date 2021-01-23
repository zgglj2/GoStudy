package slice

import (
	"fmt"
	"strings"
)

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

	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}

	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// 两个玩家轮流打上 X 和 O
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	var s2 []int
	printSlice(s2)

	s2 = append(s2, 0)
	printSlice(s2)

	s2 = append(s2, 1)
	printSlice(s2)

	s2 = append(s2, 2, 3, 4)
	printSlice(s2)
}

func printSlice(x []int){
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}