package main

import "fmt"


func main() {
	s1_from := []int{1, 2, 3}
	s1_to := make([]int, 10)

	n := copy(s1_to, s1_from)

	fmt.Println(s1_to)
	fmt.Printf("Copied %d elements\n", n)

	sl3 := []int{1, 2, 3}
	fmt.Printf("sl3 addr: %p, len: %d, cap:%d\n", &sl3, len(sl3), cap(sl3))
	sl3 = append(sl3, 4, 5, 6)
	fmt.Println(sl3)
	fmt.Printf("sl3 addr: %p, len: %d, cap:%d\n", &sl3, len(sl3), cap(sl3))
}
