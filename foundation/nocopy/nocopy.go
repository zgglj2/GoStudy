package main

import (
	"fmt"
	"sync"
)

type noCopy struct {
}

func (*noCopy) Lock()   {}
func (*noCopy) Unlock() {}

type App struct {
	noCopy noCopy
	sync.Mutex
	Name string `json:"name"`
	Type string `json:"type"`
}

func (a *App) Login() error {
	fmt.Println("app login")
	return nil
}

type Mydemo struct {
	App
	sync.Mutex
	Myversion string
}

func (a *Mydemo) demo() {
	a.Login()
}

func main() {
	mux := sync.Mutex{}
	myapp := App{
		Mutex: mux,
		Name:  "demoapp",
		Type:  "v1",
	}
	mydemo := Mydemo{
		App:   myapp,
		Mutex: mux,
	}
	mydemo2 := Mydemo{
		App:       myapp,
		Mutex:     mux,
		Myversion: "demopp",
	}
	mydemo.demo()
	mydemo2.Login()
	mydemo2.demo()
}
