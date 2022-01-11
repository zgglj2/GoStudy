package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func simpleGet() {
	client := &http.Client{
		Timeout:   time.Second * 10,
		Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
	}

	resp, err := client.Get("https://127.0.0.1:9090/")
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
	apiUrl := "https://127.0.0.1:9090/get"

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

	client := &http.Client{
		Timeout:   time.Second * 10,
		Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
	}
	resp, err := client.Get(u.String())
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
	url := "https://127.0.0.1:9090/post"
	// 表单数据
	contentType := "application/x-www-form-urlencoded"
	data := "name=小王子&age=18"
	// json
	// contentType := "application/json"
	// data := `{"name":"小王子","age":18}`
	client := &http.Client{
		Timeout:   time.Second * 10,
		Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
	}
	resp, err := client.Post(url, contentType, strings.NewReader(data))
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
