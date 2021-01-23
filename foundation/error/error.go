package error

import (
	"fmt"
	"time"
	"math"
)

//内建接口
//type error interface {
//	Error() string
//}

type DivideError struct {
	devidee int
	devider int
}

func (de *DivideError) Error() string{
	strFormat := `
	Can not devide 0
	devidee: %d
	devider: 0
`
	return fmt.Sprintf(strFormat, de.devidee)
}

func Divide(devidee int, devider int) (result int, error string) {
	if devider == 0 {
		data := DivideError{
			devidee:devidee,
			devider:devider,
		}
		error = data.Error()
		return
	} else {
		return devidee / devider, ""
	}
}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	return math.Sqrt(x), nil
}

func main() {
	if result, err := Divide(100, 10); err == "" {
		fmt.Println("100/10 = ", result)
	}
	if _, err := Divide(100, 0); err != "" {
		fmt.Println("error is ", err)
	}

	if err := run(); err != nil {
		fmt.Println(err)
	}

	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
