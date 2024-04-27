package function

import (
	"strconv"
)

func SplitFlags(str string, result []string, i int, newStr, spaceStr string) ([]string, int, string, string) {
	flags := []string{"(hex)", "(bin)", "(cap)", "(up)", "(low)", "(cap, ", "(up, ", "(low, "}
	stf := ""
	index := ""
	ch := 0
	ch1 := 0
	k := 0
	notValid := false
	k, ch, ch1, stf = CheckFlag(i, str, stf, k, ch, ch1)
	for _, c := range flags {
		notValid = false
		if stf == c && ch == 1 {
			if newStr != "" {
				result = append(result, newStr)
				newStr = ""
			}
			spaceStr, newStr, i = SpacesFlag(str, newStr, spaceStr, i, ch, ch1, k)
			result, spaceStr = ModificationByFlags(result, spaceStr, stf)
			return result, i + len(stf), newStr, spaceStr
		}
		if stf == c && ch == 2 {
			for j := k; j < len(str); j++ {
				k++
				if str[j] == ')' {
					break
				}
				index += string(str[j])
			}
			for j := 0; j < len(index); j++ {
				if !IsNum(rune(index[j])) {
					notValid = true
					break
				}
			}
			if notValid {
				break
			}
			spaceStr, newStr, i = SpacesFlag(str, newStr, spaceStr, i, ch, k, ch1)
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
