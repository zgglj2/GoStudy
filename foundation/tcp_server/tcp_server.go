package main

import (
	"flag"
	"fmt"
	"net"
)

func checkError(error error, info string) {
	if error != nil {
		panic("ERROR: " + info + " " + error.Error()) // terminate program
	}
}

func initServer(hostPort string) *net.TCPListener {
	serverAddr, err := net.ResolveTCPAddr("tcp", hostPort)
	checkError(err, "Resolving address:port failed: '"+hostPort+"'")
	listener, err := net.ListenTCP("tcp", serverAddr)
	checkError(err, "ListenTCP: ")
	println("Listening to: ", listener.Addr().String())
	return listener
}

func main() {
	flag.Parse()
	if flag.NArg() != 2 {
		panic("usage: host port")
	}
	hostAndPort := fmt.Sprintf("%s:%s", flag.Arg(0), flag.Arg(1))
	listener := initServer(hostAndPort)

}
