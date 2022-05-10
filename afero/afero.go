package main

import "github.com/spf13/afero"

func main() {
	// fs := new(afero.MemMapFs)
	// f, err := afero.TempFile(fs,"", "ioutil-test")
	fs := afero.NewMemMapFs()
	afs := &afero.Afero{Fs: fs}
	f, err := afs.TempFile("", "ioutil-test")
	if err != nil {
		panic(err)
	}
	f.WriteString("Hello, World!")

}
