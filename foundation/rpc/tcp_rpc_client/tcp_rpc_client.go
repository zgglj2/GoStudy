package main

import (
	"fmt"
	"log"
	"net/rpc"
	"time"
)

type Args struct {
	A, B int
}
type Quotient struct {
	Quo, Rem int
}

func main() {

	for {
		client, err := rpc.Dial("tcp", "127.0.0.1:1234")
		if err != nil {
			log.Fatal("dialing: ", err)
		}
		args := Args{17, 8}
		var reply int
		err = client.Call("Arith.Multiply", &args, &reply)
		if err != nil {
			client.Close()
			log.Fatal("Arith.Multiply error: ", err)
		}
		fmt.Printf("Arith.Multiply: %d*%d=%d\n", args.A, args.B, reply)

		var quot Quotient
		err = client.Call("Arith.Devide", &args, &quot)
		if err != nil {
			client.Close()
			log.Fatal("Arith.Devide error: ", err)
		}
		fmt.Printf("Arith.Devide: %d/%d=%d, remain %d\n", args.A, args.B, quot.Quo, quot.Rem)

		client.Close()
		time.Sleep(time.Second * 5)
	}

}
