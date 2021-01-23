package pointer

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // 指向 i
	fmt.Println(*p) // 通过指针读取 i 的值
	*p = 21         // 通过指针设置 i 的值
	fmt.Println(i)  // 查看 i 的值

	p = &j         // 指向 j
	*p = *p / 37   // 通过指针对 j 进行除法运算
	fmt.Println(j) // 查看 j 的值

	var i1 = 5
	fmt.Printf("An integer: %d, it's location in memory: %p\n", i1, &i1)

	var intP *int
	intP = &i1
	fmt.Printf("The value at memory location %p is %d\n", intP, *intP)

	s := "good bye"
	var ps *string = &s
	*ps = "ciao"
	fmt.Printf("Here is the pointer ps: %p\n", ps) // prints address
	fmt.Printf("Here is the string *ps: %s\n", *ps) // prints string
	fmt.Printf("Here is the string s: %s\n", s) // prints same string

	//const ii = 5
	//ptr2 := &ii //error: cannot take the address of i
	//ptr3 := &10 //error: cannot take the address of 10

	//var pp *int = nil
	//*pp = 0 //invalid memory address or nil pointer dereference
}
