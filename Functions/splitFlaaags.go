package function

import (
	"strconv"
)

func SplitFlags(str string, result []string, i int, newStr, spaceStr string) ([]string, int, string, string) {
	flags := []string{"(hex)", "(bin)", "(cap)", "(up)", "(low)", "(cap, ", "(up, ", "(low, "}
	stf := ""
	index := ""
	paramOrNot := 0
	endFlag := 0
	endFlagParam := 0
	lenValid := true
	notValid := false
	endFlagParam, paramOrNot, endFlag, stf = CheckFlag(i, str, stf, endFlagParam, paramOrNot, endFlag)
	for _, c := range flags {
		notValid = false
		if stf == c && paramOrNot == 1 {
			if newStr != "" {
				result = append(result, newStr)
				newStr = ""
			}
			result, spaceStr, newStr, i = SpacesFlag(result, str, newStr, spaceStr, i, paramOrNot, endFlag, endFlagParam)
			result = ModificationByFlags(result, stf)
			return result, i + len(stf), newStr, spaceStr
		}
		if stf == c && paramOrNot == 2 {
			for j := endFlagParam; j < len(str); j++ {
				endFlagParam++
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
			result, spaceStr, newStr, i = SpacesFlag(result, str, newStr, spaceStr, i, paramOrNot, endFlag, endFlagParam)
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
