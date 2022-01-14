package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("strings.HasPrefix(\"Hello.bak\", \"He\"): ", strings.HasPrefix("Hello.bak", "He"))
	fmt.Println("strings.HasSuffix(\"Hello.bak\", \"ak\"): ", strings.HasSuffix("Hello.bak", "ak"))

	fmt.Println("strings.Contains(\"Hello.bak\", \"lo\"): ", strings.Contains("Hello.bak", "lo"))

	var str string = "Hi, I'm Marc, Hi."

	fmt.Printf("The position of \"Marc\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "Marc"))

	fmt.Printf("The position of the first instance of \"Hi\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "Hi"))
	fmt.Printf("The position of the last instance of \"Hi\" is: ")
	fmt.Printf("%d\n", strings.LastIndex(str, "Hi"))

	fmt.Printf("The position of \"Burger\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "Burger"))

	fmt.Println("strings.Replace(\"Marc\", \"glj\"): ", strings.Replace(str, "Marc", "glj", -1))

	str = "Hello, how is it going, Hugo?"
	var manyG = "gggggggggg"

	fmt.Printf("Number of H's in %s is: ", str)
	fmt.Printf("%d\n", strings.Count(str, "H"))

	fmt.Printf("Number of double g's in %s is: ", manyG)
	fmt.Printf("%d\n", strings.Count(manyG, "gg"))

	var origS string = "Hi there! "

	newS := strings.Repeat(origS, 3)
	fmt.Printf("The new repeated string is: %s\n", newS)

	var orig string = "Hey, how are you George?"
	var lower string
	var upper string

	fmt.Printf("The original string is: %s\n", orig)
	lower = strings.ToLower(orig)
	fmt.Printf("The lowercase string is: %s\n", lower)
	upper = strings.ToUpper(orig)
	fmt.Printf("The uppercase string is: %s\n", upper)

	fmt.Println("strings.TrimSpace(\"  Hello.bak\"): ", strings.TrimSpace("      Hello.bak "))
	fmt.Println("strings.TrimSpace(\"  Hello.bak\"): ", strings.Trim("      Hello.bak ", " "))
	fmt.Println("strings.TrimSpace(\"  Hello.bak\"): ", strings.TrimLeft("      Hello.bak ", " "))
	fmt.Println("strings.TrimSpace(\"  Hello.bak\"): ", strings.TrimRight("      Hello.bak ", " "))

	str = "The quick brown fox jumps over the lazy dog"
	sl := strings.Fields(str)
	fmt.Printf("Splitted in slice: %v\n", sl)
	for _, val := range sl {
		fmt.Printf("%s - ", val)
	}
	fmt.Println()
	str2 := "GO1|The ABC of Go|25"
	sl2 := strings.Split(str2, "|")
	fmt.Printf("Splitted in slice: %v\n", sl2)
	for _, val := range sl2 {
		fmt.Printf("%s - ", val)
	}
	fmt.Println()
	str3 := strings.Join(sl2, ";")
	fmt.Printf("sl2 joined by ;: %s\n", str3)

	reader := strings.NewReader("hello world")
	fmt.Println(reader.ReadByte())
	fmt.Println(reader.ReadRune())
	fmt.Println(reader.ReadRune())
	fmt.Println(reader.ReadRune())

	var origs string = "666"
	var an int
	var newSs string

	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)

	an, _ = strconv.Atoi(origs)
	fmt.Printf("The integer is: %d\n", an)
	an = an + 5
	newSs = strconv.Itoa(an)
	fmt.Printf("The new string is: %s\n", newSs)

	var str4 string = "Hi, I'm Marc, Marc, Hi."
	fmt.Println("strings.Replace(\"Marc\", \"glj\"): ", strings.Replace(str4, "Marc", "glj", 1))
}
