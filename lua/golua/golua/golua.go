package main

import "github.com/aarzilli/golua/lua"

func adder(L *lua.State) int {
	a := (int64)(L.ToInteger(1))
	b := (int64)(L.ToInteger(2))
	L.PushInteger(a + b)
	return 1 // number of return values
}

func main() {
	L := lua.NewState()
	L.OpenLibs()
	defer L.Close()

	// push "print" function on the stack
	L.GetGlobal("print")
	// push the string "Hello World!" on the stack
	L.PushString("Hello World-1")
	L.PushString("123")
	// call print with two argument, expecting no results
	L.Call(2, 0)

	// executes a string of lua code
	if err := L.DoString(`print("Hello World-2")`); err != nil {
		panic(err)
	}

	// executes a file
	if err := L.DoFile("hello.lua"); err != nil {
		panic(err)
	}

	// excutes go function in lua
	L.Register("adder", adder)
	L.DoString("print(adder(2, 2))")

}
