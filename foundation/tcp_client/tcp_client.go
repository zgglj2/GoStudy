package main

import (
	"flag"
	"fmt"
	"net"

	"GoStudy/foundation/proto"
)

func main() {
	flag.Parse()
	host := "127.0.0.1"
	port := "8080"
	if flag.NArg() == 2 {
		host = flag.Arg(0)
		port = flag.Arg(1)
	}
	hostAndPort := fmt.Sprintf("%s:%s", host, port)
	conn, err := net.Dial("tcp", hostAndPort)
	if err != nil {
		fmt.Println("dial fail, err: ", err)
		return
	}
	defer conn.Close()

	// reader := bufio.NewReader(os.Stdin)
	// for {
	// 	fmt.Print("input: ")
	// 	msg, err := reader.ReadString('\n')
	// 	if err != nil {
	// 		fmt.Println("read input failed, err: ", err)
	// 		continue
	// 	}
	// 	msg = strings.TrimSpace(msg)
	// 	conn.Write([]byte(msg))
	// }
	for i := 0; i < 200; i++ {
		msg := "hello, hello, how are you?"
		b, err := proto.Encode(msg)
		if err != nil {
			fmt.Println("proto encode failed, err: ", err)
			continue
		}
		conn.Write(b)
	}
}
