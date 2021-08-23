package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fpath := "/root/test/test.abc"
	fmt.Println(filepath.Base(fpath))
	fmt.Println(filepath.Dir(fpath))
	fmt.Println(filepath.Ext(fpath))

}
