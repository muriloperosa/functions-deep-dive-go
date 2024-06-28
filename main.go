package main

import (
	"fmt"

	"github.com/muriloperosa/go-function-deep-dive/anonymous"
	"github.com/muriloperosa/go-function-deep-dive/functionsasvalueandtypes"
	"github.com/muriloperosa/go-function-deep-dive/recursion"
	"github.com/muriloperosa/go-function-deep-dive/variadic"
)

func main() {
	fmt.Println("== Anonymous ==")
	anonymous.Anonymous()

	fmt.Println("== Functions As Value And Types ==")
	functionsasvalueandtypes.FunctionsAsValueAndTypes()

	fmt.Println("== Recursion ==")
	recursion.Recursion()

	fmt.Println("== Variadic ==")
	variadic.Variadic()
}
