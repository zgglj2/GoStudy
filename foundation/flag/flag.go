package main

import (
	"flag"
	"os"
)

func main() {
	NewLine := flag.Bool("n", false, "print newline")

	flag.PrintDefaults()
	flag.Parse()
	var s string
	for i := 0; i < flag.NArg(); i++ {
		if i > 0 {
			s += " "
			if *NewLine {
				s += "\n"
			}
		}
		s += flag.Arg(i)
	}
	os.Stdout.WriteString(s)
}
