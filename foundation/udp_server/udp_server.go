package main

import (
	"flag"
	"fmt"
	"net"
)

func main() {
	flag.Parse()
	host := "127.0.0.1"
	port := "8080"
	if flag.NArg() == 2 {
		host = flag.Arg(0)
		port = flag.Arg(1)
	}
	// portInt, err := strconv.Atoi(port)
	// if err != nil {
	// 	fmt.Println("port is invalid, err: ", err)
	// 	return
	// }

	// listener, err := net.ListenUDP("udp", &net.UDPAddr{
	// 	IP:   net.IPv4(0, 0, 0, 0),
	// 	Port: portInt,
	// })

	addr, err := net.ResolveUDPAddr("udp", host+":"+port)
	if err != nil {
		fmt.Println("resolv failed, err: ", err)
		return
	}
	listener, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("listen failed, err: ", err)
		return
	}

	defer listener.Close()

	var data [1024]byte
	for {
		n, addr, err := listener.ReadFromUDP(data[:])
		if err != nil {
			fmt.Println("read from udp failed, err: ", err)
			continue
		}
		fmt.Printf("data: %s, addr: %v, count: %d\n", string(data[:n]), addr, n)
		_, err = listener.WriteToUDP(data[:n], addr)
		if err != nil {
			fmt.Println("write to udp failed, err: ", err)
			continue
		}

	}
}
