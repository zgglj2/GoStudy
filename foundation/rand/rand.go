package rand

import "fmt"
import "math/rand"
import "time"

func main() {
	rand.Seed(time.Now().Unix())
	fmt.Println("Rand number1: ", rand.Intn(10))
	fmt.Println("Rand number2: ", rand.Intn(10))
}
