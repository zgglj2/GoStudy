package main

import (
	"math"
	"fmt"
)

type Vertex struct {
	x, y float64
}

func (v Vertex)Abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func (v *Vertex)Scale(f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func Scale(v *Vertex, f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

func main() {
	v := Vertex{3 ,4}
	fmt.Println(v.Abs())
	fmt.Println(Abs(v))
	v.Scale(10)
	fmt.Println(v)
	Scale(&v, 10)
	fmt.Println(v)
}
