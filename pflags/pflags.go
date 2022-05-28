package main

import (
	"fmt"
	"time"

	flag "github.com/spf13/pflag"
)

func main() {
	username := flag.StringP("username", "u", "", "用户名,默认为空")
	password := flag.StringP("password", "p", "", "密码,默认为空")
	host := flag.StringP("host", "h", "127.0.0.1", "主机名,默认 127.0.0.1")
	port := flag.IntP("port", "P", 3306, "端口号,默认为空")
	// eg: 1s  100ms  1.5h  1d
	period := flag.DurationP("period", "i", 1*time.Second, "sleep period")

	flag.Parse()
	fmt.Printf("Sleeping for %v...\n", *period)
	time.Sleep(*period)

	fmt.Printf("username=%v password=%v host=%v port=%v\n", *username, *password, *host, *port)
}
