package main

import (
	"github.com/aarzilli/golua/lua"
	"github.com/ambrevar/golua/unicode"
)

func main() {
	L := lua.NewState()
	defer L.Close()
	L.OpenLibs()
	unicode.GoLuaReplaceFuncs(L)

	L.DoString(`print(string.len("résumé"))`)
	L.DoString(`print(string.upper("résumé"))`)
	L.DoString(`print(string.gsub("résumé", "\\pL", "[$0]"))`)
}
