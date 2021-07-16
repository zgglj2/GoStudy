package main

import (
	"fmt"

	"github.com/mitchellh/go-ps"
)

func main() {
	p, err := ps.Processes()
	if err != nil {
		fmt.Printf("err: %s", err)
		return
	}

	if len(p) <= 0 {
		fmt.Println("should have processes")
		return
	}

	for _, p1 := range p {

		fmt.Println(p1.Executable())
		fmt.Println(p1)
		// break
	}

}
