package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/axgle/mahonia"
)

func CmdExec(cmd string) (string, error) {
	var c *exec.Cmd
	var data string
	system := runtime.GOOS
	if system == "windows" {
		argArray := strings.Split("/c "+cmd, " ")
		c = exec.Command("cmd", argArray...)
	} else {
		c = exec.Command("/bin/sh", "-c", cmd)
	}
	out, err := c.CombinedOutput()
	if err != nil {
		return data, err
	}
	data = string(out)
	if system == "windows" {
		dec := mahonia.NewDecoder("gbk")
		data = dec.ConvertString(data)
	}
	return data, nil
}

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

	str, err := CmdExec("dir")
	if err != nil {
		fmt.Printf("CmdExec error: ", err) //
		os.Exit(1)
	}
	fmt.Println("dir: ", str)
}
