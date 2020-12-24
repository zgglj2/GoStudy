package main

import "fmt"

func main() {
	mm := map[int]string{1:"a",2:"b",3:"c"}
	fmt.Println(mm)

	mm[2] = mm[2] + "2"
	fmt.Println(mm)

	mm[4] = ""
	fmt.Println(mm)

	e, ok := mm[5]
	fmt.Println(e, ok)

	delete(mm, 4)
	fmt.Println(mm)

}
