package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func simpleGet() {
	resp, err := http.Get("http://127.0.0.1:9090/")
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read from resp.Body failed, err:%v\n", err)
		return
	}
	fmt.Println(string(body))
}

func getWithArgs() {
	apiUrl := "http://127.0.0.1:9090/get"

	u, err := url.ParseRequestURI(apiUrl)
	if err != nil {
		fmt.Printf("parse url requestUrl failed, err:%v\n", err)
	}

	// URL param
	data := url.Values{}
	data.Set("name", "小王子")
	data.Set("age", "18")
	u.RawQuery = data.Encode() // URL encode

	fmt.Println(u.String())
	resp, err := http.Get(u.String())
	if err != nil {
		fmt.Printf("post failed, err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get resp failed, err:%v\n", err)
		return
	}
	fmt.Println(string(b))
}

func postWithArgs() {
	url := "http://127.0.0.1:9090/post"
	// 表单数据
	contentType := "application/x-www-form-urlencoded"
	data := "name=小王子&age=18"
	// json
	// contentType := "application/json"
	// data := `{"name":"小王子","age":18}`
	resp, err := http.Post(url, contentType, strings.NewReader(data))
	if err != nil {
		fmt.Printf("post failed, err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get resp failed, err:%v\n", err)
		return
	}
	fmt.Println(string(b))
}
func main() {
	simpleGet()

	getWithArgs()

	postWithArgs()
}
