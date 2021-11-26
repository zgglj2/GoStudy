package main

import (
	"fmt"
	"net"
	"strings"
)

func GetOutboundIP() (ip string, err error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	// fmt.Println(localAddr.String())
	ip = strings.Split(localAddr.String(), ":")[0]
	return
}

func main() {
	ip, err := GetOutboundIP()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("ip: ", ip)
}
