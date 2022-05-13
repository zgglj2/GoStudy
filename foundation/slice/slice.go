package main

import (
	"fmt"
	"strings"
	"sync"
	"unsafe"
)

func sliceModify(s []int) {
	s[0] = 100
}

func sliceAppend(s []int) []int {
	s = append(s, 100)
	return s
}

func sliceAppendPtr(s *[]int) {
	*s = append(*s, 100)
}

// 注意：Go语言中所有的传参都是值传递（传值），都是一个副本，一个拷贝。
// 拷贝的内容是非引用类型（int、string、struct等这些），在函数中就无法修改原内容数据；
// 拷贝的内容是引用类型（interface、指针、map、slice、chan等这些），这样就可以修改原内容数据。
func SliceFn() {
	fmt.Println("-----------------------------------------------------")
	// 参数为引用类型slice：外层slice的len/cap不会改变，指向的底层数组会改变
	s := []int{1, 1, 1}
	newS := sliceAppend(s)
	// 函数内发生了扩容
	fmt.Println(s, len(s), cap(s))
	// [1 1 1] 3 3
	fmt.Println(newS, len(newS), cap(newS))
	// [1 1 1 100] 4 6

	s2 := make([]int, 0, 5)
	newS = sliceAppend(s2)
	// 函数内未发生扩容
	fmt.Println(s2, s2[0:5], len(s2), cap(s2))
	// [] [100 0 0 0 0] 0 5
	fmt.Println(newS, newS[0:5], len(newS), cap(newS))
	// [100] [100 0 0 0 0] 1 5

	// 参数为引用类型slice的指针：外层slice的len/cap会改变，指向的底层数组会改变
	sliceAppendPtr(&s)
	fmt.Println(s, len(s), cap(s))
	// [1 1 1 100] 4 6
	sliceModify(s)
	fmt.Println(s, len(s), cap(s))
	// [100 1 1 100] 4 6
	fmt.Println("-----------------------------------------------------")
}
func SliceEmptyOrNil() {
	var slice1 []int
	// slice1 is nil slice
	slice2 := make([]int, 0)
	// slcie2 is empty slice
	var slice3 = make([]int, 2)
	// slice3 is zero slice
	if slice1 == nil {
		fmt.Println("slice1 is nil.")
		// 会输出这行
	}
	if slice2 == nil {
		fmt.Println("slice2 is nil.")
		// 不会输出这行
	}
	fmt.Println(slice3) // [0 0]
}
func SliceConcurrencySafeByMutex() {
	var lock sync.Mutex //互斥锁
	a := make([]int, 0)
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			lock.Lock()
			defer lock.Unlock()
			a = append(a, i)
		}(i)
	}
	wg.Wait()
	fmt.Println(len(a))
	// equal 10000
}

func SliceConcurrencySafeByChanel() {
	buffer := make(chan int)
	a := make([]int, 0)
	// 消费者
	go func() {
		for v := range buffer {
			a = append(a, v)
		}
	}()
	// 生产者
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			buffer <- i
		}(i)
	}
	wg.Wait()
	fmt.Println(len(a))
	// equal 10000
}

func SliceInit() {
	// 初始化方式1：直接声明
	var slice1 []int
	fmt.Println(len(slice1), cap(slice1))
	// 0, 0
	slice1 = append(slice1, 1)
	fmt.Println(len(slice1), cap(slice1))
	// 1, 1

	// 初始化方式2：使用字面量
	slice2 := []int{1, 2, 3, 4}
	fmt.Println(len(slice2), cap(slice2))
	// 4, 4

	// 初始化方式3：使用make创建slice
	slice3 := make([]int, 3, 5)
	// make([]T, len, cap) cap不传则和len一样
	fmt.Println(len(slice3), cap(slice3))
	// 3, 5
	fmt.Println(slice3[0], slice3[1], slice3[2])
	// 0, 0, 0
	// fmt.Println(slice3[3], slice3[4])
	// panic: runtime error: index out of range [3] with length 3
	slice3 = append(slice3, 1)
	fmt.Println(len(slice3), cap(slice3))
	// 4, 5

	// 初始化方式4: 从切片或数组“截取”
	arr := [100]int{}
	for i := range arr {
		arr[i] = i
	}
	slcie4 := arr[1:3]
	slice5 := make([]int, len(slcie4))
	copy(slice5, slcie4)
	fmt.Println(len(slcie4), cap(slcie4), unsafe.Sizeof(slcie4))
	// 2，99，24
	fmt.Println(len(slice5), cap(slice5), unsafe.Sizeof(slice5))
	// 2，2，24
}

func SliceGrowing() {
	slice1 := []int{}
	for i := 0; i < 10; i++ {
		slice1 = append(slice1, i)
		fmt.Println(len(slice1), cap(slice1))
	}
}

func main() {
	var numbers1 = make([]int, 3, 5)

	printSlice(numbers1)

	var numbers2 []int

	printSlice(numbers2)

	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(numbers)

	fmt.Println("numbers ==", numbers)

	fmt.Println("numbers[1:4] ==", numbers[1:4])

	fmt.Println("numbers[:3] ==", numbers[:3])

	fmt.Println("numbers[4:] ==", numbers[4:])

	numbers4 := make([]int, 0, 5)
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

	numbers5 = append(numbers5, 2, 3, 4)
	printSlice(numbers5)

	numbers6 := make([]int, len(numbers5), (cap(numbers5))*2)

	copy(numbers6, numbers5)
	printSlice(numbers6)

	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}

	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
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

	SliceFn()
	SliceEmptyOrNil()
	SliceConcurrencySafeByMutex()
	SliceConcurrencySafeByChanel()
	SliceInit()
	SliceGrowing()
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
