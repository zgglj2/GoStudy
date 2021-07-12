package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4}
	fmt.Printf("s: %v, addr: %p, len: %d, cap:%d\n", s, &s, len(s), cap(s))
	s2 := append(s[:1], s[2:]...)
	fmt.Printf("s: %v, addr: %p, len: %d, cap:%d\n", s2, &s2, len(s2), cap(s2))
	x, s3 := s[len(s)-1], s[:len(s)-1]
	fmt.Printf("x: %v\n", x)
	fmt.Printf("s: %v, addr: %p, len: %d, cap:%d\n", s3, &s3, len(s3), cap(s3))
}
