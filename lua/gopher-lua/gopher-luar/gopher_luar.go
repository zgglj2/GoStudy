package main

import (
	"fmt"
	"strconv"

	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
)

type User struct {
	Name  string
	token string
}

func (u *User) SetToken(t string) {
	u.token = t
}

func (u *User) Token() string {
	return u.token
}

func Example_basic() {
	const script = `
	print("Hello from Lua, " .. u.Name .. "!")
	u:SetToken("12345")
	`

	L := lua.NewState()
	defer L.Close()

	u := &User{
		Name: "Tim",
	}
	L.SetGlobal("u", luar.New(L, u))
	if err := L.DoString(script); err != nil {
		panic(err)
	}

	fmt.Println("Lua set your token to:", u.Token())
	// Output:
	// Hello from Lua, Tim!
	// Lua set your token to: 12345
}

func Example() {
	const test = `
for i = 1, 3 do
		print(msg, i)
end
print(user)
print(user.Name, user.Age)
`

	type person struct {
		Name string
		Age  int
	}

	L := lua.NewState()
	defer L.Close()

	user := &person{"Dolly", 46}

	L.SetGlobal("print", luar.New(L, fmt.Println))
	L.SetGlobal("msg", luar.New(L, "foo"))
	L.SetGlobal("user", luar.New(L, user))

	L.DoString(test)
	// Output:
	// foo 1
	// foo 2
	// foo 3
	// &{Dolly 46}
	// Dolly 46
}

type Ref struct {
	Index  int
	Number *int
	Title  *string
}

func (r *Ref) GetNumber() int {
	return *r.Number
}

func (r *Ref) GetTitle() string {
	return *r.Title
}

// Pointers to structs and structs within pointers are automatically dereferenced.
func Example_pointers() {
	const test = `
	local t = newRef(10, 'foo')
	print(t.Index, t.Number, t.Title)   -- 17 0xc0000aa8f8 0xc000175c40
	print(t:GetNumber())                -- 10
	print(t:GetTitle())                 -- foo
	`
	newRef := func(number int, title string) *Ref {
		n := new(int)
		*n = number
		t := new(string)
		*t = title
		return &Ref{Index: 17, Number: n, Title: t}
	}

	L := lua.NewState()
	defer L.Close()
	L.SetGlobal("print", luar.New(L, fmt.Println))
	L.SetGlobal("newRef", luar.New(L, newRef))

	L.DoString(test)
	// Output:
	// 17 10 foo
}

// Slices must be looped over with 'ipairs'.
func Example_slices() {
	const test = `
	print(type(names))
	print(names)
	for i, v in ipairs(names) do
		print(i, v)
	end
	`
	L := lua.NewState()
	defer L.Close()

	names := []string{"alfred", "alice", "bob", "frodo"}

	L.SetGlobal("print", luar.New(L, fmt.Println))
	L.SetGlobal("names", luar.New(L, names))

	L.DoString(test)
	// Output:
	// 1 alfred
	// 2 alice
	// 3 bob
	// 4 frodo
}

func ExampleInit() {
	const code = `
-- Lua tables auto-convert to slices in Go-function calls.
local res = foo {10, 20, 30, 40}

-- The result is a map-proxy.
print(res['1'], res['2'])

-- Which we may explicitly convert to a table.
res = luar.unproxify(res)
for k,v in pairs(res) do
	print(k,v)
end
`

	foo := func(args []int) (res map[string]int) {
		res = make(map[string]int)
		for i, val := range args {
			res[strconv.Itoa(i)] = val * val
		}
		return
	}

	L := lua.NewState()
	defer L.Close()

	L.SetGlobal("print", luar.New(L, fmt.Println))
	L.SetGlobal("foo", luar.New(L, foo))

	res := L.DoString(code)
	if res != nil {
		fmt.Println("Error:", res)
	}
	// Output:
	// 400 900
	// 1 400
	// 0 100
	// 3 1600
	// 2 900
}

func main() {
	Example_basic()
	fmt.Println("-------------------")
	Example()
	fmt.Println("-------------------")
	Example_pointers()
	fmt.Println("-------------------")
	Example_slices()
	fmt.Println("-------------------")
	ExampleInit()
	fmt.Println("-------------------")
}
