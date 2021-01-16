// math.go 文件

// smath 提供一些简单的数学函数
package godoc

import (
	"errors"
)

// Math 其实啥都没做
type Math struct{
	simplenumber int
}

// Error 没有做啥的error
var (
	ErrorSimple = errors.New("simple err")
	ErrorNotSimple = errors.New("not simple err")
)

// New 创建一个Math对象
func New()*Math{
	return &Math{}
}

// Add 两数相加
//
//     result := Add(1,2)
//     result = 3
//
// 欢迎使用
func Add(n1,n2 int)(result int){
	return n1+n2
}

// BadAdd 两数相加
//
// BUG(zhaojun) 明显加错了
func BadAdd(n1,n2 int)(result int){
	return n1
}

// Deprecated: OldAdd 老旧的方法，不建议使用了
func OldAdd(n1,n2 int)(int){
	return n1+n2
}

// Add 两数相加
func (m *Math)Add(n1,n2 int)int{
	return n1+n2
}

// add 非外部方法
func add(n1,n2 int) int {
	return n1+n2
}
