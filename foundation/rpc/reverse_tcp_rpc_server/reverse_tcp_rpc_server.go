package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

type Args struct {
	A, B int
}
type Quotient struct {
	Quo, Rem int
}

func main() {
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}

	clientChan := make(chan *rpc.Client)

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Fatal("Accecpt error: ", err)
			}
			clientChan <- rpc.NewClient(conn)
		}
	}()

	for {
		client := <-clientChan

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
	}

}
