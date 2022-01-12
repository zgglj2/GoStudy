package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	searchIn := "John: 2578.34 William: 4567.23 Steve: 5632.18"
	pat := "[0-9]+.[0-9]+" //正则

	f := func(s string) string {
		v, _ := strconv.ParseFloat(s, 32)
		return strconv.FormatFloat(v*2, 'f', 2, 32)
	}

	if ok, _ := regexp.Match(pat, []byte(searchIn)); ok {
		fmt.Println("Match found")
	}

	re, _ := regexp.Compile(pat)
	str := re.ReplaceAllString(searchIn, "##.#")
	fmt.Println(str)

	str2 := re.ReplaceAllStringFunc(searchIn, f)
	fmt.Println(str2)

	info := "users:((\"java\",24487,36))"
	reg := regexp.MustCompile(`users:\(\("(.*?)",(.*?),(.*?)\)`)
	r := reg.FindSubmatch([]byte(info))
	for index, item := range r {
		fmt.Printf("item[%d]: %s\n", index, string(item))
	}
	r2 := reg.FindStringSubmatch(info)
	for index, item := range r2 {
		fmt.Printf("item[%d]: %s\n", index, string(item))
	}
}
