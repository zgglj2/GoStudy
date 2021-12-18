package main

import (
	"context"
	"fmt"

	"github.com/olivere/elastic/v7"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Married bool   `json:"married"`
}

func main() {
	client, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"))
	if err != nil {
		fmt.Printf("new elastic client failed, err:%v\n", err)
		return
	}
	fmt.Println("connect to es success")
	p1 := Person{Name: "rion", Age: 22, Married: false}
	put1, err := client.Index().Index("user").BodyJson(p1).Do(context.Background())
	if err != nil {
		fmt.Printf("index failed, err:%v\n", err)
		return
	}
	fmt.Printf("Indexed user %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)

}
