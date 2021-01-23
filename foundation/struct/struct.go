package _struct

import "fmt"

type Vertex struct {
	X, Y int
}

func main() {
	fmt.Println(Vertex{1, 2})

	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v)

	p := &v
	p.X = 1e9
	fmt.Println(v)

	var (
		v1 = Vertex{1, 2}  // 创建一个 Vertex 类型的结构体
		v2 = Vertex{X: 1}  // Y:0 被隐式地赋予
		v3 = Vertex{}      // X:0 Y:0
		p2  = &Vertex{1, 2} // 创建一个 *Vertex 类型的结构体（指针）
	)
	fmt.Println(v1, p2, v2, v3)
}

