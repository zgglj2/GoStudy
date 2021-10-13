package main

import (
	"fmt"
	"os"
	"runtime"
)

// func GetCpuInfo() string {
// 	var size uint32 = 128
// 	var buffer = make([]uint16, size)
// 	var index = uint32(copy(buffer, syscall.StringToUTF16("Num:")) - 1)
// 	nums := syscall.StringToUTF16Ptr("NUMBER_OF_PROCESSORS")
// 	arch := syscall.StringToUTF16Ptr("PROCESSOR_ARCHITECTURE")
// 	r, err := syscall.GetEnvironmentVariable(nums, &buffer[index], size-index)
// 	if err != nil {
// 		return ""
// 	}
// 	index += r
// 	index += uint32(copy(buffer[index:], syscall.StringToUTF16(" Arch:")) - 1)
// 	r, err = syscall.GetEnvironmentVariable(arch, &buffer[index], size-index)
// 	if err != nil {
// 		return syscall.UTF16ToString(buffer[:index])
// 	}
// 	index += r
// 	return syscall.UTF16ToString(buffer[:index+r])
// }

func main() {
	fmt.Println("os: ", runtime.GOOS)
	fmt.Println("os: ", runtime.GOARCH)
	fmt.Println("os: ", runtime.GOROOT())
	fmt.Println("os: ", runtime.NumCPU())
	fmt.Println("os: ", runtime.NumGoroutine())

	where := func() {
		pc, file, line, ok := runtime.Caller(1)
		fmt.Printf("%s:%d\n", file, line)
		fmt.Printf("%v:%v", pc, ok)
	}
	where()
	fmt.Println()
	// if runtime.GOOS == "windows" {
	// 	fmt.Println(GetCpuInfo())
	// }
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}
