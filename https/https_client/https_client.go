package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	// pool := x509.NewCertPool()
	// caCertPath := "../ca.crt"

	// caCrt, err := ioutil.ReadFile(caCertPath)
	// if err != nil {
	// 	fmt.Println("ReadFile err:", err)
	// 	return
	// }
	// pool.AppendCertsFromPEM(caCrt)

	// tr := &http.Transport{
	// 	TLSClientConfig: &tls.Config{RootCAs: pool},
	// }
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://127.0.0.1:8081")
	if err != nil {
		fmt.Println("Get error:", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ReadAll error:", err)
		return
	}
	fmt.Println(string(body))
}
