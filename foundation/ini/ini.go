package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}
type RedisConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Password string `ini:"password"`
	Database string `ini:"database"`
	Test     string `ini:"test"`
}

type Config struct {
	MysqlConfig `ini:"mysql"`
	RedisConfig `ini:"redis"`
}

func loadIni(filename string, data interface{}) error {
	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)
	fmt.Println("Data type: ", t.Kind())
	if t.Kind() != reflect.Ptr {
		return fmt.Errorf("data type is not ptr")
	}
	fmt.Println("Elem type: ", t.Elem().Kind())
	if t.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("data elem type is not struct")
	}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Open file(%s) fail\n", filename)
		return err
	}
	defer file.Close()
	inputReader := bufio.NewReader(file)
	configName := ""
	for {
		line, err := inputReader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				if len(line) == 0 {
					break
				}
			} else {
				return fmt.Errorf("read file failed, err: %s", err)
			}
		}

		fmt.Printf("The input was: %s", line)
		line = strings.TrimSpace(line)
		lineLen := len(line)
		if lineLen < 1 {
			fmt.Println("  Empty line")
			continue
		}
		if line[0] == '#' || line[0] == ';' {
			continue
		}
		if line[0] == '[' {
			if lineLen < 2 || line[lineLen-1] != ']' {
				return fmt.Errorf("section line is not valid")
			}
			sectionName := line[1 : lineLen-1]
			sectionName = strings.TrimSpace(sectionName)
			if len(sectionName) < 1 {
				return fmt.Errorf("empty section name")
			}
			fmt.Println("  Section name: ", sectionName)
			fmt.Println(t.Elem().NumField())
			for i := 0; i < t.Elem().NumField(); i++ {

				if t.Elem().Field(i).Tag.Get("ini") == sectionName {
					configName = t.Elem().Field(i).Name
					fmt.Println(configName)
					break
				}
			}
		} else {
			if configName == "" {
				return fmt.Errorf("config line not under section line")
			}
			if lineLen < 2 || !strings.Contains(line, "=") {
				return fmt.Errorf("config line is not valid")
			}
			if len(line) < 2 {
				return fmt.Errorf("empty config")
			}
			idx := strings.Index(line, "=")
			key := strings.TrimSpace(line[0:idx])
			value := strings.TrimSpace(line[idx+1:])
			if len(key) < 1 {
				return fmt.Errorf("key is empty")
			}
			fmt.Println("  key: ", key)
			fmt.Println("  value: ", value)
			sValue := v.Elem().FieldByName(configName)
			sType := sValue.Type()
			fmt.Println(sType)
			keyName := ""
			for i := 0; i < sType.NumField(); i++ {
				if sType.Field(i).Tag.Get("ini") == key {
					keyName = sType.Field(i).Name
					fmt.Println(keyName)
					break
				}
			}

			if keyName == "" {
				fmt.Printf("config line key(%s) not in struct, drop it\n", key)
				continue
			}
			fileObj := sValue.FieldByName(keyName)
			switch fileObj.Type().Kind() {
			case reflect.String:
				fileObj.SetString(value)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				iValue, err := strconv.ParseInt(value, 10, 64)
				if err != nil {
					return err
				}
				fileObj.SetInt(iValue)
			case reflect.Bool:
				iValue, err := strconv.ParseBool(value)
				if err != nil {
					return err
				}
				fileObj.SetBool(iValue)
			case reflect.Float32, reflect.Float64:
				iValue, err := strconv.ParseFloat(value, 64)
				if err != nil {
					return err
				}
				fileObj.SetFloat(iValue)
			}

		}

	}
	fmt.Println("")
	return nil
}

func main() {
	var config Config
	err := loadIni("./conf.ini", &config)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	fmt.Printf("%#v", config)
}
