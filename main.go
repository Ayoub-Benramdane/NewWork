package main

import (
	"fmt"
	function "function/Functions"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	if os.Args[1] != "sample.txt" || os.Args[2] != "result.txt" {
		return
	}
	sample, _ := os.ReadFile(os.Args[1])
	result := function.Split(string(sample))
	fmt.Println(result)
	resFinal := os.WriteFile(os.Args[2], []byte(strings.Trim(function.PrintStrFinal(result), " ")), 0644)
	if resFinal != nil {
		return
	}
}
