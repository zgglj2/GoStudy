package main

import (
	"compress/gzip"
	"fmt"
	"os"
)

func main() {
	fw, err := os.Create("demo.gzip")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fw.Close()

	gw := gzip.NewWriter(fw)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer gw.Close()

	content := "Hello, test gzip\nNew line test"
	gw.Header.Name = "demo.txt"

	_, err = gw.Write([]byte(content))
	if err != nil {
		fmt.Println(err)
		return
	}

	fr, err := os.Open("demo.gzip")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fr.Close()

	gr, err := gzip.NewReader(fr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer gr.Close()

	buf := make([]byte, 1024*1024*10)

	_, err = gr.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(buf)

}
