package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{21, 3, 55, 4, 8, 100, 67}
	fmt.Println(sort.SearchInts(a, 55))
	sort.Ints(a)
	fmt.Println(a)
	fmt.Println(sort.SearchInts(a, 55))
}
