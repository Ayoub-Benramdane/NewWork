package function

func SplitPonc(result []string, newStr, str string, i int) ([]string, string, int) {
	strPonc := ""
	if newStr != "" {
		result = append(result, newStr)
		newStr = ""
	}
	for j := i; j < len(str); j++ {
		if str[j] != ' ' && !IsPonc(rune(str[j])) {
			break
		}
		i++
		if str[j] != ' ' && IsPonc(rune(str[j])) {
			strPonc += string(str[j])
		}
	}
	if i < len(str) {
		result[len(result)-1] += strPonc + " "
	} else {
		result[len(result)-1] += strPonc
	}
	return result, newStr, i
}
