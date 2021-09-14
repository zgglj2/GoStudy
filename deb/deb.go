package main

import (
	"fmt"

	"github.com/isbm/go-deb"
)

func main() {
	pkg, err := deb.OpenPackageFile("gzip_1.6-5+b1_amd64.deb", false)
	if err != nil {
		panic(err)
	}

	// fmt.Printf("Loaded package: %v - %s\n", p, p.Summary())
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
