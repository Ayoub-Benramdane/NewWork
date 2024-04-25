package main

import (
	"fmt"
	function "function/Functions"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	sample, _ := os.ReadFile(os.Args[1])
	args := function.SplitWhiteSpaces(string(sample))
	fmt.Println(args)
	// resFinal := os.WriteFile(os.Args[2], []byte(strings.Trim(function.PrintStr(args), " ")), 0644)
	// if resFinal != nil {
	// 	return
	// }
}
