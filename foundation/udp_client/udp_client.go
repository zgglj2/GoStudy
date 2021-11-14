package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	flag.Parse()
	host := "127.0.0.1"
	port := "8080"
	if flag.NArg() == 2 {
		host = flag.Arg(0)
		port = flag.Arg(1)
	}

	addr, err := net.ResolveUDPAddr("udp", host+":"+port)
	if err != nil {
		fmt.Println("resolv failed, err: ", err)
		return
	}
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println("dial udp failed, err: ", err)
		return
	}

	defer conn.Close()

	var reply [1024]byte
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("input: ")
		msg, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("read input failed, err: ", err)
			continue
		}
		msg = strings.TrimSpace(msg)
		_, err = conn.Write([]byte(msg))
		if err != nil {
			fmt.Println("write to udp failed, err: ", err)
			continue
		}
		n, addr, err := conn.ReadFromUDP(reply[:])
		if err != nil {
			fmt.Println("read from udp failed, err: ", err)
			continue
		}
		fmt.Printf("data: %s, addr: %v, count: %d\n", string(reply[:n]), addr, n)
	}

}
