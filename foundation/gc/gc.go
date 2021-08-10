package main

import (
	"fmt"
	"runtime"
	"time"
)

type Vertex struct {
	x, y float64
}

func final(v *Vertex) {
	fmt.Println(v)
}

func entry() {
	v := Vertex{3, 4}
	runtime.SetFinalizer(&v, final)
}

func main() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Printf("%d Kb\n", m.Alloc)

	entry()

	runtime.GC()
	time.Sleep(time.Second)

	runtime.ReadMemStats(&m)
	fmt.Printf("%d Kb\n", m.Alloc)

}
