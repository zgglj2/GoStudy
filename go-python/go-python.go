package main

import (
	"fmt"
	"log"
	"os"

	python3 "github.com/jshiwam/cpy3x/pycore"
)

func main() {
	python3.Py_Initialize()
	defer python3.Py_Finalize()

	cwd, err := os.Getwd()
	if err != nil {
		log.Panic(err)
	}
	code := python3.PyRun_SimpleString(fmt.Sprintf(`
import sys
sys.path.append(r"%s")
`, cwd))
	if code != 0 {
		log.Panic(fmt.Errorf("call `PyRun_SimpleString` error"))
	}
	f, err := os.Create("test.py")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()
	if _, err := f.WriteString(`
print("test")
	`); err != nil {
		panic(err)
	}
	code, err = python3.Py_Main([]string{"python", "test.py"})
	if err != nil {
		panic(err)
	}
	if code != 0 {
		log.Panic(fmt.Errorf("call `Py_Main` error"))
	}
}
