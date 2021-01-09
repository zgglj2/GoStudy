package main

import "fmt"

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

func main() {
	if result, error := Divide(100, 10); error == "" {
		fmt.Println("100/10 = ", result)
	}
	if _, error := Divide(100, 0); error != "" {
		fmt.Println("error is ", error)
	}
}
