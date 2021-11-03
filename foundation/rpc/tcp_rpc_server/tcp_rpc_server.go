package main

import (
	"errors"
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

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Devide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("devide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main() {
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}

	arith := new(Arith)
	rpc.Register(arith)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accecpt error: ", err)
		}
		go func(net.Conn) {
			rpc.ServeConn(conn)
			conn.Close()
		}(conn)
	}
}
