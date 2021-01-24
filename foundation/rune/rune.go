package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var str = "hello 你好"
	fmt.Println("len(str): ", len(str))
	fmt.Println("len([]rune(str)): ", len([]rune(str)))
	fmt.Println("utf8.RuneCountInString(str): ", utf8.RuneCountInString(str))

	s1 := "白萝卜"
	s2 := []rune(s1)

	s2[0] = '红'
	fmt.Println(string(s2))

	c1 := "红"
	c2 := '红'
	fmt.Printf("c1:%T, c2:%T\n", c1, c2)

}
