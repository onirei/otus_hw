package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	inputStr := "Hello, OTUS!"
	outputStr := stringutil.Reverse(inputStr)
	fmt.Println(outputStr)
}
