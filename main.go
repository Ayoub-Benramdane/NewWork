package main

import (
	function "function/Functions"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	sample, _ := os.ReadFile(os.Args[1])
	args := function.SplitWhiteSpaces(string(sample))
	resFinal := os.WriteFile(os.Args[2], []byte(function.PrintStrFinal(args)), 0644)
	if resFinal != nil {
		return
	}
}
