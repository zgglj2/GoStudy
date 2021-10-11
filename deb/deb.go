package main

import (
	"archive/tar"
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
	"strings"

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
		fmt.Printf("%v %v %v %v %5v %v %v %v %v %v\n",
			fi.Mode(),
			fi.Mode().Perm(),
			fi.Owner(),
			fi.Group(),
			fi.Size(),
			fi.ModTime().UTC().Format("Jan 02 15:04"),
			fi.Name(),
			fi.Digest(),
			fi.Mode().IsDir(),
			fi.Mode().IsRegular())
	}

	reader, err := deb.PayloadReader("gzip_1.6-5+b1_amd64.deb", false)
	if err != nil {
		panic(err)
	}
	var databuf bytes.Buffer
	for {
		header, err := reader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		fmt.Println(header.Name, header.Size)
		if header.Typeflag == tar.TypeReg && header.Mode&0111 != 0 && header.Size != 0 {
			if strings.HasPrefix(header.Name, "/usr/share/doc") {
				continue
			}

			databuf.Reset()
			_, err = io.Copy(&databuf, reader)
			if err != nil {
				fmt.Println(err)
				continue
			}
			has := md5.Sum(databuf.Bytes())
			md5str := fmt.Sprintf("%x", has)
			fmt.Println(header.Name, md5str)
		}
	}
}
