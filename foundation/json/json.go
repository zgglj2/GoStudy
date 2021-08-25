package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Address struct {
	Type    string
	City    string
	Country string
}
type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

func main() {
	pa := &Address{"private", "Aartselaar", "Belgium"}
	wa := &Address{"work", "Boom", "Belgium"}
	vc := VCard{"Jan", "Kersschot", []*Address{pa, wa}, "none"}
	fmt.Println(vc)
	for i := 0; i < len(vc.Addresses); i++ {
		fmt.Println(vc.Addresses[i])
	}
	js, _ := json.Marshal(vc)
	fmt.Printf("type: %T\n", js)
	fmt.Printf("JSON format: %s\n", js)
	fmt.Println(string(js))

	file, _ := os.OpenFile("vcard.json", os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()

	enc := json.NewEncoder(file)
	err := enc.Encode(vc)
	if err != nil {
		fmt.Println(err)
	}

}
