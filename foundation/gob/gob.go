package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"os"
)

type P struct {
	X, Y, Z int
	Name    string
}
type Q struct {
	X, Y *int32
	Name string
}

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
	var network bytes.Buffer        // Stand-in for a network connection
	enc := gob.NewEncoder(&network) // Will write to network.
	dec := gob.NewDecoder(&network) // Will read from network.
	// Encode (send) the value.
	err := enc.Encode(P{3, 4, 5, "Pythagoras"})
	if err != nil {
		fmt.Println("encode error:", err)
		return
	}
	// Decode (receive) the value.
	var q Q
	err = dec.Decode(&q)
	if err != nil {
		fmt.Println("decode error:", err)
		return
	}
	fmt.Printf("%q: {%d,%d}\n", q.Name, *q.X, *q.Y)

	pa := &Address{"private", "Aartselaar", "Belgium"}
	wa := &Address{"work", "Boom", "Belgium"}
	vc := VCard{"Jan", "Kersschot", []*Address{pa, wa}, "none"}
	// fmt.Printf("%v: \n", vc) // {Jan Kersschot [0x126d2b80 0x126d2be0] none}:
	// using an encoder:
	file, _ := os.OpenFile("vcard.gob", os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	enc = gob.NewEncoder(file)
	err = enc.Encode(vc)
	if err != nil {
		fmt.Println("Error in encoding gob")
		return
	}
}
