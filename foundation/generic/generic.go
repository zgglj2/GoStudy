package main

import "fmt"

func SumIntsOrFloats[K comparable, V int | float64](m map[K]V) V {
	var sum V
	for _, v := range m {
		sum += v
	}
	return sum
}
func main() {

	ints := map[string]int{
		"first":  34,
		"second": 12,
	}
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}
	fmt.Printf("泛型计算结果, Ints结果:%v, Floats结果:%v\n",
		SumIntsOrFloats(ints),
		SumIntsOrFloats(floats))
}
