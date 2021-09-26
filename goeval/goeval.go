package main

import (
	"fmt"

	"github.com/PaulXu-cn/goeval"
)

func main() {
	if re, err := goeval.Eval(
		"",
		"fmt.Print(\"Hello World!\")",
		"fmt"); nil == err {
		fmt.Print(string(re))
	} else {
		fmt.Print(err.Error())
	}
}
