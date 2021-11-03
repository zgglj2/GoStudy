package main

import (
	"fmt"
	"os"

	"github.com/bwmarrin/snowflake"
)

func main() {
	n, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for i := 0; i < 3; i++ {
		id := n.Generate()
		fmt.Println("id", id)
		fmt.Println(
			"node: ", id.Node(),
			"step: ", id.Step(),
			"time: ", id.Time(),
		)
		idbyte, _ := id.MarshalJSON()
		fmt.Println(string(idbyte))
	}
}
