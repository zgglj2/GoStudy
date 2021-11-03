package main

import (
	"errors"
	"log"
	"net"
	"net/rpc"
	"time"
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
	arith := new(Arith)
	rpc.Register(arith)

	for {
		conn, err := net.Dial("tcp", "127.0.0.1:1234")
		if err != nil {
			log.Fatal("dialing: ", err)
		}

		rpc.ServeConn(conn)
		conn.Close()
		time.Sleep(time.Second * 5)
	}

}
