package function

import (
	"fmt"
	"strconv"
)

func SplitFlags(str string, result []string, i int, newStr, spaceStr string) ([]string, int, string, string) {
	flags := []string{"(hex)", "(bin)", "(cap)", "(up)", "(low)", "(cap, ", "(up, ", "(low, "}
	stf := ""
	index := ""
	ch := 0
	ch1 := 0
	k := 0
	lenValid := true
	notValid := false
	k, ch, ch1, stf = CheckFlag(i, str, stf, k, ch, ch1)
	for _, c := range flags {
		notValid = false
		if stf == c && ch == 1 {
			if newStr != "" {
				result = append(result, newStr)
				newStr = ""
			}
			fmt.Println(spaceStr)
			result, spaceStr, newStr, i = SpacesFlag(result, str, newStr, spaceStr, i, ch, ch1, k)
			result = ModificationByFlags(result, stf)
			return result, i + len(stf), newStr, spaceStr
		}
		if stf == c && ch == 2 {
			for j := k; j < len(str); j++ {
				k++
				if str[j] == ')' {
					lenValid = true
					break
				} else {
					lenValid = false
				}
				index += string(str[j])
			}
			for j := 0; j < len(index); j++ {
				if !IsNum(rune(index[j])) {
					notValid = true
					break
				}
			}
			if notValid || !lenValid {
				break
			}
			result, spaceStr, newStr, i = SpacesFlag(result, str, newStr, spaceStr, i, ch, ch1, k)
			if newStr != "" {
				result = append(result, newStr)
				newStr = ""
			}
			indexFinal, _ := strconv.Atoi(index)
			result = ModificationByParamsFlags(result, stf, indexFinal)
			return result, i + len(stf) + len(index) + 1, newStr, spaceStr
		}
	}
	return result, i, newStr, spaceStr
}
