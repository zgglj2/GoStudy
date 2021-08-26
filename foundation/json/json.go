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
		panic(err)
	}

	b := []byte(`{"Name": "Wednesday", "Age": 6, "Parents": ["Gomez", "Morticia"]}`)
	var f interface{}
	err = json.Unmarshal(b, &f)
	if err != nil {
		panic(err)
	}

	fmt.Println(f)

	m := f.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Printf("%s is of a type(%T) I don’t know how to handle\n", k, v)
		}
	}
}
