package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type TagType struct {
	field1 bool   "An important answer"
	field2 string "The name of the thing"
	field3 int    "How much there are"
}

func refTag(tt TagType, ix int) {
	tType := reflect.TypeOf(tt)
	tField := tType.Field(ix)
	fmt.Printf("%#v\n", tField)
}
func main() {
	tt := TagType{true, "Barak Obama", 1}
	for i := 0; i < 3; i++ {
		refTag(tt, i)
	}
	fmt.Printf("field1 size: %d\n", unsafe.Sizeof(tt.field1))
	fmt.Printf("field2 size: %d\n", unsafe.Sizeof(tt.field2))
	fmt.Printf("field3 size: %d\n", unsafe.Sizeof(tt.field3))
	fmt.Printf("struct size: %d\n", unsafe.Sizeof(tt))

}
