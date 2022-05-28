package main

import (
	"fmt"
	"time"

	flag "github.com/spf13/pflag"
)

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

type celsiusFlag struct{ Celsius }

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit) // no error check needed
	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

func main() {
	username := flag.StringP("username", "u", "", "用户名,默认为空")
	password := flag.StringP("password", "p", "", "密码,默认为空")
	host := flag.StringP("host", "h", "127.0.0.1", "主机名,默认 127.0.0.1")
	port := flag.IntP("port", "P", 3306, "端口号,默认为空")
	// eg: 1s  100ms  1.5h  1d
	period := flag.DurationP("period", "i", 1*time.Second, "sleep period")

	flag.Parse()
	fmt.Printf("Sleeping for %v...\n", *period)
	time.Sleep(*period)

	fmt.Printf("username=%v password=%v host=%v port=%v\n", *username, *password, *host, *port)
}
