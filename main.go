package main

import (
	function "function/Functions"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	sample, _ := os.ReadFile(os.Args[1])
	result := function.Split(string(sample))
	resFinal := os.WriteFile(os.Args[2], []byte(strings.Trim(function.PrintStrFinal(result), " ")), 0644)
	if resFinal != nil {
		return
	}
}
