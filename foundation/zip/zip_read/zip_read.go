package main

import (
	"archive/zip"
	"log"
)

func main() {
	zipReader, err := zip.OpenReader("../output.zip")
	if err != nil {
		panic(err)
	}
	defer zipReader.Close()

	for _, file := range zipReader.Reader.File {
		// zippedFile, err := file.Open()
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// defer zippedFile.Close()

		if file.FileInfo().IsDir() {
			log.Println("Directory:", file.Name)

		} else {
			log.Println("File:", file.Name)

			// buf, err := ioutil.ReadAll(zippedFile)
			// if err != nil {
			// 	fmt.Println(err)
			// 	return
			// }
			// fmt.Println(string(buf))
		}
	}
}
