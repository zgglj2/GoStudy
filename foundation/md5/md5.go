package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
)

func md5V(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func md5V2(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

func md5V3(str string) string {
	w := md5.New()
	io.WriteString(w, str)
	md5str := fmt.Sprintf("%x", w.Sum(nil))
	return md5str
}

func main() {
	str := "MD5testing"
	fmt.Println(md5V(str))
	fmt.Println(md5V2(str))
	fmt.Println(md5V3(str))
}
