package main

import (
	"fmt"
	"net/http"
)

func handler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Hi, This is an example of https service in golang!\n")
	fmt.Fprintf(res,
		`[{"Name":"jason","Age":35,"Weight":60.3,"Speciality":"computer science","Hobby":["tennis","swimming","reading"],"Score":725.5,"Secret":"SRRMb3ZlFFlvdSE="}]`)
}

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServeTLS(":8081", "../server.crt", "../server.key", nil)
	fmt.Println(err)
}
