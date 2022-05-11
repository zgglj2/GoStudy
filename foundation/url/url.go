package main

import (
	"fmt"
	"net/url"
)

func main() {
	apiUrl := "https://127.0.0.1:9090/get?name=小王子&age=18"

	u, err := url.ParseRequestURI(apiUrl)
	if err != nil {
		fmt.Printf("parse url requestUrl failed, err:%v\n", err)
	}
	fmt.Printf("u: %#v\n", u)
	fmt.Printf("name: %s\n", u.Query().Get("name"))
	fmt.Printf("age: %s\n", u.Query().Get("age"))
}
