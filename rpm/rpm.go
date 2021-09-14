package main

import (
	"fmt"
	"os"

	"github.com/cavaliercoder/go-rpm"
)

func ExampleMD5Check(rpm_file string) {
	// open a rpm package for reading
	f, err := os.Open(rpm_file)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	// validate md5 checksum
	if err := rpm.MD5Check(f); err == nil {
		fmt.Printf("Package passed checksum validation\n")
	} else if err == rpm.ErrMD5ValidationFailed {
		fmt.Printf("Package failed checksum validation\n")
	} else {
		panic(err)
	}

	// Output: Package passed checksum validation
}

func ExamplePackageFile_Files(rpm_file string) {
	// open a package file
	pkg, err := rpm.OpenPackageFile(rpm_file)
	if err != nil {
		panic(err)
	}

	// list each file
	files := pkg.Files()
	fmt.Printf("total %v\n", len(files))
	for _, fi := range files {
		fmt.Printf("%v %v %v %5v %v %v %v\n",
			fi.Mode().Perm(),
			fi.Owner(),
			fi.Group(),
			fi.Size(),
			fi.ModTime().UTC().Format("Jan 02 15:04"),
			fi.Name(),
			fi.Digest())
	}
}

func main() {
	rpm_file := "net-tools-1.60-78.el5.x86_64.rpm"
	ExampleMD5Check(rpm_file)
	ExamplePackageFile_Files(rpm_file)
}
