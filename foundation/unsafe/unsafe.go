package main

import (
	"fmt"
	"unsafe"
)

type File struct {
	fd   int    // 文件描述符
	name string // 文件名
}

func main() {
	fmt.Println(unsafe.Sizeof(File{}))

}
