package main

import (
	"fmt"
	"os"
)

func main() {
	env := os.Environ()
	procAttr := &os.ProcAttr{
		Env: env,
		Files: []*os.File{
			os.Stdin,
			os.Stdout,
			os.Stderr,
		},
	}

	// 1st example: list files

	// pid, err := os.StartProcess("/bin/ls", []string{"ls", "-l"}, procAttr)
	pid, err := os.StartProcess("C:\\Windows\\System32\\netstat.exe", []string{"netstat.exe", "-an"}, procAttr)
	if err != nil {
		fmt.Printf("Error %v starting process!", err) //
		os.Exit(1)
	}
	fmt.Printf("The process id is %v", pid)

}
