package main

import "fmt"

// 赋值一个常量时，之后没赋值的常量都会应用上一行的赋值表达式
const (
	a = iota // a = 0
	b        // b = 1
	c        // c = 2
	d = 5    // d = 5
	e        // e = 5
)

// 赋值两个常量，iota 只会增长一次，而不会因为使用了两次就增长两次
const (
	Apple, Banana     = iota + 1, iota + 2 // Apple=1 Banana=2
	Cherimoya, Durian                      // Cherimoya=2 Durian=3
	Elderberry, Fig                        // Elderberry=3, Fig=4

)

// 使用 iota 结合 位运算 表示资源状态的使用案例
const (
	Open    = 1 << iota // 0001
	Close               // 0010
	Pending             // 0100
)

const (
	_  = iota             // 使用 _ 忽略不需要的 iota
	KB = 1 << (10 * iota) // 1 << (10*1)
	MB                    // 1 << (10*2)
	GB                    // 1 << (10*3)
	TB                    // 1 << (10*4)
	PB                    // 1 << (10*5)
	EB                    // 1 << (10*6)
	//ZB                             // 1 << (10*7)
	//YB                             // 1 << (10*8)
)

func main() {
	fmt.Println(a, b, c, d, e)
	fmt.Println(Apple, Banana, Cherimoya, Durian, Elderberry, Fig)
	fmt.Println(Open, Close, Pending)
	fmt.Print(KB, MB, GB, TB, PB, EB)

}
