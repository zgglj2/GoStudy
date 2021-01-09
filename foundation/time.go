package main

import "fmt"
import "time"

func main() {
	fmt.Println("Now time: ", time.Now())
	fmt.Println("Now weekday: ", time.Now().Weekday())
	fmt.Println("Sunday: ", time.Sunday)
}
