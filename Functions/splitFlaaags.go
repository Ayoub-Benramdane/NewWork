package function

import (
	"strconv"
)

func SplitFlags(str string, result []string, i int) ([]string, int) {
	flags := []string{"(hex)", "(bin)", "(cap)", "(up)", "(low)", "(cap, ", "(up, ", "(low, "}
	stf := ""
	index := ""
	ch := 0
	k := 0
	for j := i; j < len(str); j++ {
		stf += string(str[j])
		if str[j] == ')' {
			ch = 1
			break
		}
		if str[j] == ' ' {
			ch = 2
			k = j + 1
			break
		}
	}
	for _, c := range flags {
		notValid := false
		if stf == c && ch == 1 {
			result = ModificationByFlags(result, stf)
			return result, i + len(stf)
		}
		if stf == c && ch == 2 {
			for j := k; j < len(str); j++ {
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
			indexFinal, _ := strconv.Atoi(index)
			result = ModificationByParamsFlags(result, stf, indexFinal)
			return result, i + len(stf) + len(index) + 1
		}
	}
	return result, i
}
