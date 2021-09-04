package main

import (
	"flag"
	"fmt"
	"net"
	"strconv"
	"syscall"
)

const maxRead = 25

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

func handleMSG(readn int, err error, buf []byte) {
	if readn > 0 {
		print("<", readn, ":")
		for i := 0; ; i++ {
			if buf[i] == 0 {
				break
			}
			fmt.Printf("%c", buf[i])
		}
		print(">")
	}
}
func connectionHandler(conn net.Conn) {
	connFrom := conn.RemoteAddr()
	fmt.Println("Connection from: ", connFrom)

	welcome := "Welcome!!!"
	wrote, err := conn.Write([]byte(welcome))
	checkError(err, "Write: wrote "+strconv.Itoa(wrote)+" bytes.")
	for {
		buf := make([]byte, maxRead+1)
		readn, err := conn.Read(buf[0:maxRead])
		buf[readn] = 0
		switch err {
		case nil:
			handleMSG(readn, err, buf)
		case syscall.Errno(0xb):
			continue
		default:
			goto DISCONNECT
		}
	}

DISCONNECT:
	err = conn.Close()
	println("Closed connection: ", connFrom)
	checkError(err, "Close: ")

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
	listener := initServer(hostAndPort)

	for {
		conn, err := listener.Accept()
		checkError(err, "Accept: ")
		go connectionHandler(conn)
	}

}
