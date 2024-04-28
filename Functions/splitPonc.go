package function

func StrPonc(str, strPonc string, i int) (string, int) {
	for j := i; j < len(str); j++ {
		if str[j] != ' ' && !IsPonc(rune(str[j])) {
			break
		}
		i++
		if str[j] != ' ' && IsPonc(rune(str[j])) {
			strPonc += string(str[j])
		}
	}
	return strPonc, i
}

func SplitPonc(result []string, newStr, str string, i int) ([]string, string, int) {
	strPonc := ""
	if newStr != "" {
		result = append(result, newStr)
		newStr = ""
	}
	if len(result) > 0 && result[len(result)-1] != "'" {
		strPonc, i = StrPonc(str, strPonc, i)
		if i <= len(str) {
			if result[len(result)-1][0] == ' ' {
				result[len(result)-1] = strPonc + " "
			} else {
				result[len(result)-1] += strPonc + " "
			}
		}
	} else {
		strPonc, i = StrPonc(str, strPonc, i)
		result = append(result, strPonc+" ")
	}
	return result, newStr, i
}
