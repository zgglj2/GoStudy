package main

import "net/http"

func f1(w http.ResponseWriter, r *http.Request) {
	str := "你好"
	w.Write([]byte(str))
}

func main() {
	http.HandleFunc("/", f1)
	http.ListenAndServe("127.0.0.1:9090", nil)
}
