package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
)

var memprofile = flag.String("memprofile", "", "write memory profile to this file")

func main() {
	flag.Parse()
	testmem := new(int)
	fmt.Println(testmem)
	// CallToFunctionWhichAllocatesLotsOfMemory()
	if *memprofile != "" {
		fmt.Println(*memprofile)
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.WriteHeapProfile(f)
		f.Close()
		return
	}
}
