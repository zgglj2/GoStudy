package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type xmldas struct {
	XMLName  xml.Name       `xml:"das"`
	DataPort string         `xml:"DataPort,attr"`
	Desc     string         `xml:"desc,attr"`
	Src      xmlsource      `xml:"source"`
	Dest     xmldestination `xml:"destination"`
}

type xmlsource struct {
	Path  string `xml:"path,attr"`
	Param string `xml:"param,attr"`
}

type xmldestination struct {
	Path  string `xml:"path,attr"`
	Param string `xml:"param,attr"`
}

func main() {
	xmlString := `<?xml version="1.0" encoding="UTF-8"?>
	<das DataPort="8250" desc="123">
		<source path="123" param="456"></source>
		<destination path="789" param="000"></destination>
	</das>`
	var v xmldas
	err := xml.Unmarshal([]byte(xmlString), &v)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%#v\n", v)

	output, err := xml.MarshalIndent(v, "", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	os.Stdout.Write([]byte(xml.Header))
	os.Stdout.Write(output)
}
