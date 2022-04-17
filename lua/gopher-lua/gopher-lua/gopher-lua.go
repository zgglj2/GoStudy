package main

import (
	"GoStudy/lua/gopher-lua/gopher-lua/mymodule"
	"bufio"
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	lua "github.com/yuin/gopher-lua"
	"github.com/yuin/gopher-lua/parse"
)

type Person struct {
	Name string
}

const luaPersonTypeName = "person"

// Registers my person type to given L.
func registerPersonType(L *lua.LState) {
	mt := L.NewTypeMetatable(luaPersonTypeName)
	L.SetGlobal("person", mt)
	// static attributes
	L.SetField(mt, "new", L.NewFunction(newPerson))
	// methods
	L.SetField(mt, "__index", L.SetFuncs(L.NewTable(), personMethods))
}

// Constructor
func newPerson(L *lua.LState) int {
	person := &Person{L.CheckString(1)}
	ud := L.NewUserData()
	ud.Value = person
	L.SetMetatable(ud, L.GetTypeMetatable(luaPersonTypeName))
	L.Push(ud)
	return 1
}

// Checks whether the first lua argument is a *LUserData with *Person and returns this *Person.
func checkPerson(L *lua.LState) *Person {
	ud := L.CheckUserData(1)
	if v, ok := ud.Value.(*Person); ok {
		return v
	}
	L.ArgError(1, "person expected")
	return nil
}

var personMethods = map[string]lua.LGFunction{
	"name": personGetSetName,
}

// Getter and setter for the Person#Name
func personGetSetName(L *lua.LState) int {
	p := checkPerson(L)
	if L.GetTop() == 2 {
		p.Name = L.CheckString(2)
		return 0
	}
	L.Push(lua.LString(p.Name))
	return 1
}

// CompileLua reads the passed lua file from disk and compiles it.
func CompileLua(filePath string) (*lua.FunctionProto, error) {
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	reader := bufio.NewReader(file)
	chunk, err := parse.Parse(reader, filePath)
	if err != nil {
		return nil, err
	}
	proto, err := lua.Compile(chunk, filePath)
	if err != nil {
		return nil, err
	}
	return proto, nil
}

// DoCompiledFile takes a FunctionProto, as returned by CompileLua, and runs it in the LState. It is equivalent
// to calling DoFile on the LState with the original source file.
func DoCompiledFile(L *lua.LState, proto *lua.FunctionProto) error {
	lfunc := L.NewFunctionFromProto(proto)
	L.Push(lfunc)
	return L.PCall(0, lua.MultRet, nil)
}

// Example shows how to share the compiled byte code from a lua script between multiple VMs.
func Example() {
	codeToShare, _ := CompileLua("hello.lua")
	a := lua.NewState()
	b := lua.NewState()
	c := lua.NewState()
	DoCompiledFile(a, codeToShare)
	DoCompiledFile(b, codeToShare)
	DoCompiledFile(c, codeToShare)
}

func receiver(ch, quit chan lua.LValue) {
	L := lua.NewState()
	defer L.Close()
	L.SetGlobal("ch", lua.LChannel(ch))
	L.SetGlobal("quit", lua.LChannel(quit))
	if err := L.DoString(`
    local exit = false
    while not exit do
      channel.select(
        {"|<-", ch, function(ok, v)
          if not ok then
            print("channel closed")
            exit = true
          else
            print("received:", v)
          end
        end},
        {"|<-", quit, function(ok, v)
            print("quit")
            exit = true
        end}
      )
    end
  `); err != nil {
		panic(err)
	}
}

func sender(ch, quit chan lua.LValue) {
	L := lua.NewState()
	defer L.Close()
	L.SetGlobal("ch", lua.LChannel(ch))
	L.SetGlobal("quit", lua.LChannel(quit))
	if err := L.DoString(`
    ch:send("1")
    ch:send("2")
  `); err != nil {
		panic(err)
	}
	ch <- lua.LString("3")
	quit <- lua.LTrue
}

type lStatePool struct {
	m     sync.Mutex
	saved []*lua.LState
}

func (pl *lStatePool) Get() *lua.LState {
	pl.m.Lock()
	defer pl.m.Unlock()
	n := len(pl.saved)
	if n == 0 {
		return pl.New()
	}
	x := pl.saved[n-1]
	pl.saved = pl.saved[0 : n-1]
	return x
}

func (pl *lStatePool) New() *lua.LState {
	L := lua.NewState()
	// setting the L up here.
	// load scripts, set global variables, share channels, etc...
	return L
}

func (pl *lStatePool) Put(L *lua.LState) {
	pl.m.Lock()
	defer pl.m.Unlock()
	pl.saved = append(pl.saved, L)
}

func (pl *lStatePool) Shutdown() {
	for _, L := range pl.saved {
		L.Close()
	}
}

// Global LState pool
var luaPool = &lStatePool{
	saved: make([]*lua.LState, 0, 4),
}

func MyWorker() {
	L := luaPool.Get()
	defer luaPool.Put(L)
	/* your code here */
	if err := L.DoFile("hello.lua"); err != nil {
		panic(err)
	}
}

func main() {
	L := lua.NewState(lua.Options{
		RegistrySize:     1024 * 20, // this is the initial size of the registry
		RegistryMaxSize:  1024 * 80, // this is the maximum size that the registry can grow to. If set to `0` (the default) then the registry will not auto grow
		RegistryGrowStep: 32,        // this is how much to step up the registry by each time it runs out of space. The default is `32`.

		CallStackSize:       120,  // this is the maximum callstack size of this LState
		MinimizeStackMemory: true, // Defaults to `false` if not specified. If set, the callstack will auto grow and shrink as needed up to a max of `CallStackSize`. If not set, the callstack will be fixed at `CallStackSize`.

		IncludeGoStackTrace: true,
	})
	// L := lua.NewState()
	defer L.Close()
	// lua.RegistrySize = 1024 * 20
	// lua.RegistryGrowStep = 32
	// lua.CallStackSize = 120

	// executes a string of lua code
	if err := L.DoString(`print("Hello World-2")`); err != nil {
		panic(err)
	}

	// executes a file
	if err := L.DoFile("hello.lua"); err != nil {
		panic(err)
	}

	Double := func(L *lua.LState) int {
		lv := L.ToInt(1)            /* get argument */
		L.Push(lua.LNumber(lv * 2)) /* push result */
		return 1                    /* number of results */
	}

	L.SetGlobal("double", L.NewFunction(Double)) /* Original lua_setglobal uses stack... */
	L.DoString(`print(double(2));return "123"`)

	lv := L.Get(-1) // get the value at the top of the stack
	if str, ok := lv.(lua.LString); ok {
		// lv is LString
		fmt.Println(string(str))
	}
	if lv.Type() != lua.LTString {
		fmt.Printf("string required., get %s\n", lv.Type().String())
	}

	L.DoString(`return {"banana","orange","apple"}`)
	lv = L.Get(-1) // get the value at the top of the stack
	if tbl, ok := lv.(*lua.LTable); ok {
		// lv is LTable
		fmt.Println(L.ObjLen(tbl))
	}

	L.DoString(`return nil`)
	lv = L.Get(-1)       // get the value at the top of the stack
	if lv == lua.LTrue { // correct
		fmt.Println(lv)
	}

	// if bl, ok := lv.(lua.LBool); ok { // wrong, nil is not LBool
	// 	fmt.Println(bl)
	// }

	if lua.LVIsFalse(lv) { // lv is nil or false
		fmt.Println("lv is nil nor false")
	}

	if lua.LVAsBool(lv) { // lv is neither nil nor false
		fmt.Println("lv is neither nil nor false")
	}

	// Panic := func(L *lua.LState) int {
	// 	panic("bad end")
	// 	return 1 /* number of results */
	// }
	// L.SetGlobal("panic", L.NewFunction(Panic)) /* Original lua_setglobal uses stack... */
	// L.DoString(`panic();return "123"`)
	// fmt.Println("after panic")

	L2 := lua.NewState(lua.Options{SkipOpenLibs: true})
	defer L2.Close()
	for _, pair := range []struct {
		n string
		f lua.LGFunction
	}{
		{lua.LoadLibName, lua.OpenPackage}, // Must be first
		{lua.BaseLibName, lua.OpenBase},
		{lua.TabLibName, lua.OpenTable},
	} {
		if err := L2.CallByParam(lua.P{
			Fn:      L2.NewFunction(pair.f),
			NRet:    0,
			Protect: true,
		}, lua.LString(pair.n)); err != nil {
			panic(err)
		}
	}
	if err := L.DoFile("hello.lua"); err != nil {
		panic(err)
	}

	L.PreloadModule("mymodule", mymodule.Loader)
	if err := L.DoFile("mymodule.lua"); err != nil {
		panic(err)
	}

	if err := L.DoFile("max.lua"); err != nil {
		panic(err)
	}
	if err := L.CallByParam(lua.P{
		Fn:      L.GetGlobal("max"),
		NRet:    1,
		Protect: true,
	}, lua.LNumber(10), lua.LNumber(20)); err != nil {
		panic(err)
	}
	ret := L.Get(-1) // returned value
	L.Pop(1)         // remove received value
	fmt.Println(ret)

	registerPersonType(L)
	if err := L.DoString(`
        p = person.new("Steeve")
        print(p:name()) -- "Steeve"
        p:name("Alice")
        print(p:name()) -- "Alice"
    `); err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	// set the context to our LState
	L.SetContext(ctx)
	err := L.DoString(`
		local clock = os.clock
		function sleep(n)  -- seconds
			local t0 = clock()
			while clock() - t0 <= n do end
		end
		sleep(3)
	`)
	// err.Error() contains "context deadline exceeded"
	fmt.Println("err:", err)

	ctx, cancel = context.WithCancel(context.Background())
	L.SetContext(ctx)
	defer cancel()
	L.DoString(`
		function coro()
			local i = 0
			while true do
				coroutine.yield(i)
					i = i+1
			end
			return i
		end
	`)
	co, cocancel := L.NewThread()
	defer cocancel()
	fn := L.GetGlobal("coro").(*lua.LFunction)

	_, err, values := L.Resume(co, fn) // err is nil
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("values:", values)
	_, err, values = L.Resume(co, fn)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("values:", values)

	cancel() // cancel the parent context

	_, err, values = L.Resume(co, fn) // err is NOT nil : child context was canceled
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("values:", values)
	}

	Example()

	ch := make(chan lua.LValue)
	quit := make(chan lua.LValue)
	go receiver(ch, quit)
	go sender(ch, quit)
	time.Sleep(3 * time.Second)

	defer luaPool.Shutdown()
	go MyWorker()
	go MyWorker()
	time.Sleep(3 * time.Second)
}
