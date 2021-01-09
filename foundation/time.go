package main

import "fmt"
import "time"

func main() {
	t := time.Now()
	fmt.Println("Now time: ", t)
	fmt.Println("Now weekday: ", t.Weekday())
	fmt.Println("Now Hour: ", t.Hour())
	fmt.Println("Sunday: ", time.Sunday)
}
