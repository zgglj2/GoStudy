package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"log"

	example "github.com/rpcxio/rpcx-examples"
	"github.com/smallnest/rpcx/client"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

func main() {
	flag.Parse()

	d, err := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")
	if err != nil {
		fmt.Println("NewPeer2PeerDiscovery failed, err: ", err)
		return
	}
	option := client.DefaultOption

	conf := &tls.Config{
		InsecureSkipVerify: true,
	}

	option.TLSConfig = conf

	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, option)
	defer xclient.Close()

	args := &example.Args{
		A: 10,
		B: 20,
	}

	reply := &example.Reply{}
	err = xclient.Call(context.Background(), "Mul", args, reply)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	log.Printf("%d * %d = %d", args.A, args.B, reply.C)

}
