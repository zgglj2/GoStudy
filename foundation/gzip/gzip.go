package main

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// tar + gzip
	var buf bytes.Buffer
	_ = compress("../reflect", &buf)

	// write the .tar.gzip
	fileToWrite, err := os.OpenFile("./compress.tar.gz", os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.FileMode(0600))
	if err != nil {
		panic(err)
	}
	defer fileToWrite.Close()
	if _, err := io.Copy(fileToWrite, &buf); err != nil {
		panic(err)
	}

	file, err := os.Open("./compress.tar.gz")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// untar write
	if err = uncompress(file, "./uncompressHere/"); err != nil {
		panic(err)
	}
}

func compress(src string, buf io.Writer) error {
	// tar > gzip > buf
	zr := gzip.NewWriter(buf)
	defer zr.Close()
	tw := tar.NewWriter(zr)
	defer tw.Close()

	walker := func(file string, fi os.FileInfo, err error) error {
		// generate tar header
		header, lerr := tar.FileInfoHeader(fi, file)
		if lerr != nil {
			return lerr
		}

		// must provide real name
		// (see https://golang.org/src/archive/tar/common.go?#L626)
		header.Name = filepath.ToSlash(file)

		// write header
		if err := tw.WriteHeader(header); err != nil {
			return err
		}
		// if not a dir, write file content
		if !fi.IsDir() {
			data, err := os.Open(file)
			if err != nil {
				return err
			}
			if _, err := io.Copy(tw, data); err != nil {
				return err
			}
		}
		return nil
	}
	// walk through every file in the folder
	err := filepath.Walk(src, walker)
	if err != nil {
		panic(err)
	}
	return nil
}

// check for path traversal and correct forward slashes
func validRelPath(p string) bool {
	if p == "" || strings.Contains(p, `\`) || strings.HasPrefix(p, "/") || strings.Contains(p, "../") {
		return false
	}
	return true
}

func uncompress(src io.Reader, dst string) error {
	// ungzip
	zr, err := gzip.NewReader(src)
	if err != nil {
		return err
	}
	zr.Close()

	// untar
	tr := tar.NewReader(zr)

	// uncompress each element
	for {
		header, err := tr.Next()
		if err == io.EOF {
			break // End of archive
		}
		if err != nil {
			return err
		}
		target := strings.TrimLeft(header.Name, "../")

		// validate name against path traversal
		if !validRelPath(target) {
			return fmt.Errorf("tar contained invalid name error %q\n", target)
		}

		// add dst + re-format slashes according to system
		target = filepath.Join(dst, header.Name)
		// if no join is needed, replace with ToSlash:
		// target = filepath.ToSlash(header.Name)

		// check the type
		switch header.Typeflag {

		// if its a dir and it doesn't exist create it (with 0755 permission)
		case tar.TypeDir:
			if _, err := os.Stat(target); err != nil {
				if err := os.MkdirAll(target, 0755); err != nil {
					return err
				}
			}
		// if it's a file create it (with same permission)
		case tar.TypeReg:
			fileToWrite, err := os.OpenFile(target, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))
			if err != nil {
				return err
			}
			// copy over contents
			if _, err := io.Copy(fileToWrite, tr); err != nil {
				return err
			}
			// manually close here after each file operation; defering would cause each file close
			// to wait until all operations have completed.
			fileToWrite.Close()
		}
	}
	return nil
}
