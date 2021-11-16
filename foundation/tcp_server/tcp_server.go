package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net"
	"time"

	"GoStudy/foundation/proto"
)

func process(conn net.Conn) {
	defer conn.Close()

	connFrom := conn.RemoteAddr()
	fmt.Println("Connection from: ", connFrom)

	reader := bufio.NewReader(conn)

	for {
		recvStr, err := proto.Decode(reader)
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("read from client failed, err: ", err)
			break
		}
		if recvStr == "" {
			fmt.Println(err)
			return
		}
		fmt.Println("recv: ", recvStr)
		// conn.Write(buf[0:n])
	}
}
func main() {
	flag.Parse()
	host := "127.0.0.1"
	port := "8080"
	if flag.NArg() == 2 {
		host = flag.Arg(0)
		port = flag.Arg(1)
	}
	hostAndPort := fmt.Sprintf("%s:%s", host, port)
	listener, err := net.Listen("tcp", hostAndPort)
	if err != nil {
		fmt.Println("listen failed, err: ", err)
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept failed, err: ", err)
			time.Sleep(time.Second)
			continue
		}
		go process(conn)
	}
}
