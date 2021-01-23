package channel

import "fmt"

func sum(s []int, ch chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	ch <- sum
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}
func main() {
	s := []int{1, 2, 3, 9, 6}
	ch := make(chan int)
	go sum(s[:len(s)/2], ch)
	go sum(s[len(s)/2:], ch)
	x, y := <-ch, <-ch
	fmt.Println(x, y, x+y)

	ch2 := make(chan int, 2)
	ch2 <- 1
	ch2 <- 2
	fmt.Println(<-ch2, <-ch2)

	c := make(chan int, 10)
	go fibonacci(cap(c), c)

	for i := range c {
		fmt.Println(i)
	}
}
