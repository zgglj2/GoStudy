package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Id   int
	Name string
}

func (s Student) Hello() {
	fmt.Println("我是一个学生")
}

func (s Student) EchoName(name string) {
	fmt.Println("我的名字是：", name)
}

type People struct {
	Student // 匿名字段
}

func main() {
	var name string = "非凡"

	// TypeOf会返回目标数据的类型，比如int/float/struct/指针等
	reflectType := reflect.TypeOf(name)

	// valueOf返回目标数据的的值，比如上文的"非凡"
	reflectValue := reflect.ValueOf(name)

	fmt.Println("type: ", reflectType)
	fmt.Println("value: ", reflectValue)

	fmt.Println("------------------------------------------")

	s := Student{Id: 1, Name: "非凡"}
	// 获取目标对象
	t := reflect.TypeOf(s)
	// .Name()可以获取去这个类型的名称
	fmt.Println("这个类型的名称是:", t.Name())

	// 获取目标对象的值类型
	v := reflect.ValueOf(s)
	// .NumField()来获取其包含的字段的总数
	for i := 0; i < t.NumField(); i++ {
		// 从0开始获取Student所包含的key
		key := t.Field(i)

		// 通过interface方法来获取key所对应的值
		value := v.Field(i).Interface()

		fmt.Printf("第%d个字段是：%s:%v = %v \n", i+1, key.Name, key.Type, value)
	}

	// 通过.NumMethod()来获取Student里头的方法
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("第%d个方法是：%s:%v\n", i+1, m.Name, m.Type)
	}

	fmt.Println("------------------------------------------")
	p := People{Student{Id: 1, Name: "非凡"}}

	t = reflect.TypeOf(p)
	// 这里需要加一个#号，可以把struct的详情都给打印出来
	// 会发现有Anonymous:true，说明是匿名字段
	fmt.Printf("%#v\n", t.Field(0))

	// 取出这个学生的名字的详情打印出来
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 1}))

	// 获取匿名字段的值的详情
	v = reflect.ValueOf(p)
	fmt.Printf("%#v\n", v.Field(0))

	fmt.Println("------------------------------------------")
	s = Student{Id: 1, Name: "非凡"}
	t = reflect.TypeOf(s)

	// 通过.Kind()来判断对比的值是否是struct类型
	if k := t.Kind(); k == reflect.Struct {
		fmt.Println("bingo")
	}

	num := 1
	numType := reflect.TypeOf(num)
	if k := numType.Kind(); k == reflect.Int {
		fmt.Println("bingo")
	}
	fmt.Println("------------------------------------------")
	sp := &Student{Id: 1, Name: "非凡"}

	v = reflect.ValueOf(sp)

	// 修改值必须是指针类型否则不可行
	if v.Kind() != reflect.Ptr {
		fmt.Println("不是指针类型，没法进行修改操作")
		return
	}

	// 获取指针所指向的元素
	v = v.Elem()

	// 获取目标key的Value的封装
	namep := v.FieldByName("Name")

	if namep.Kind() == reflect.String {
		namep.SetString("小学生")
	}

	fmt.Printf("%#v \n", *sp)

	// 如果是整型的话
	test := 888
	testV := reflect.ValueOf(&test)
	testV.Elem().SetInt(666)
	fmt.Println(test)

	fmt.Println("------------------------------------------")

	s = Student{Id: 1, Name: "非凡"}

	v = reflect.ValueOf(s)

	// 获取方法控制权
	// 官方解释：返回v的名为name的方法的已绑定（到v的持有值的）状态的函数形式的Value封装
	mv := v.MethodByName("EchoName")
	// 拼凑参数
	args := []reflect.Value{reflect.ValueOf("非凡")}

	// 调用函数
	mv.Call(args)
	fmt.Println("------------------------------------------")
}
