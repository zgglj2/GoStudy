package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	fmt.Println(os.Args)
	args := make([]string, 0)
	if len(os.Args) == 1 {
		args = append(args, "http://www.baidu.com")
		args = append(args, "http://www.163.com")
	} else {
		// for _, arg := range os.Args[1:] {
		// 	args = append(args, os.Args[1:]...)
		// }
		args = append(args, os.Args[1:]...)
	}
	fmt.Println(args)

	for _, url := range args {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)

	}
}
