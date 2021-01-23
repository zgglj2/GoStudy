package _interface

import "fmt"

type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("Nokia Phone")
}

type IPhone struct {
}

func (iPhone IPhone) call() {
	fmt.Println("iPhone")
}

func describe(i Phone) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func describe2(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func do(i interface{}) {
	switch v:=i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}
func main() {
	var phone Phone
	describe(phone)
	//phone.call()
	phone = new(NokiaPhone)
	describe(phone)
	phone.call()
	phone = new(IPhone)
	describe(phone)
	phone.call()

	var i interface{}
	describe2(i)

	i = 42
	describe2(i)

	i = "hello"
	describe2(i)

	var j interface{} = "hello"
	s := j.(string)
	fmt.Println(s)

	s, ok := j.(string)
	fmt.Println(s, ok)

	f, ok := j.(float64)
	fmt.Println(f, ok)

	//f = i.(float64) // 报错(panic)
	//fmt.Println(f)

	do(21)
	do("hello")
	do(true)

}
