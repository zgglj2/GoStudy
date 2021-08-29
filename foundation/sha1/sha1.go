package main

import (
	"crypto/sha1"
	"fmt"
	"io"
)

func main() {
	hasher := sha1.New()
	io.WriteString(hasher, "test")
	b := []byte{}
	fmt.Printf("Result: %x\n", hasher.Sum(b))
	fmt.Printf("Result: %d\n", hasher.Sum(b))

	hasher.Reset()
	data := []byte("We shall overcome!")
	n, err := hasher.Write(data)
	if n != len(data) || err != nil {
		fmt.Printf("Hash write error: %v / %v", n, err)
		return
	}
	checksum := hasher.Sum(b)
	fmt.Printf("Result: %x\n", checksum)
}
