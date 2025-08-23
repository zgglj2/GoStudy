package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"strings"
)

type SystemSetting struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

type Fset struct {
	SystemSettings []SystemSetting `xml:"SystemSetting"`
}

func loadXML(path string) (Fset, error) {
	file, err := os.Open(path)
	if err != nil {
		return Fset{}, err
	}
	defer file.Close()

	var data Fset
	decoder := xml.NewDecoder(file)
	err = decoder.Decode(&data)
	if err != nil {
		return Fset{}, err
	}

	return data, nil
}

func findSetting(settings []SystemSetting, name string) (SystemSetting, bool) {
	for _, setting := range settings {
		if setting.Name == name {
			return setting, true
		}
	}

	return SystemSetting{}, false
}

func compareXML(fset1, fset2 Fset) (string, string) {
	var similarities []string
	var differences []string

	if len(fset1.SystemSettings) != len(fset2.SystemSettings) {
		differences = append(differences, fmt.Sprintf("Different number of SystemSettings: %d vs %d", len(fset1.SystemSettings), len(fset2.SystemSettings)))
	}

	for _, setting1 := range fset1.SystemSettings {
		setting2, found := findSetting(fset2.SystemSettings, setting1.Name)

		if !found {
			differences = append(differences, fmt.Sprintf("SystemSetting not found in second file: %s", setting1.Name))
		} else if setting1.Value != setting2.Value {
			differences = append(differences, fmt.Sprintf("Different value for SystemSetting %s: %s vs %s", setting1.Name, setting1.Value, setting2.Value))
		} else {
			similarities = append(similarities, fmt.Sprintf("Same value for SystemSetting %s: %s", setting1.Name, setting1.Value))
		}
	}

	for _, setting2 := range fset2.SystemSettings {
		_, found := findSetting(fset1.SystemSettings, setting2.Name)

		if !found {
			differences = append(differences, fmt.Sprintf("SystemSetting not found in first file: %s", setting2.Name))
		}
	}

	return strings.Join(similarities, "\n"), strings.Join(differences, "\n")
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Please provide two XML file paths as arguments.")
		os.Exit(1)
	}

	filePath1 := os.Args[1]
	filePath2 := os.Args[2]

	fset1, err := loadXML(filePath1)
	if err != nil {
		panic(err)
	}

	fset2, err := loadXML(filePath2)
	if err != nil {
		panic(err)
	}

	similarities, differences := compareXML(fset1, fset2)
	if differences == "" {
		println("XML files are identical")
	} else {
		println("XML files are not identical:")
		println("Differences:")
		println(differences)
	}
	if similarities != "" {
		println("Similarities:")
		println(similarities)
	}
}
