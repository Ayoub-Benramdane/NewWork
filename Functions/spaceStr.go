package function

func SpaceStr(result []string, newStr, spaceStr, str string, count int) ([]string, string, int, string) {
	if count == 0 && spaceStr != "" {
		result = append(result, spaceStr)
		spaceStr = ""
	}
	if !IsPonc(rune(str[0])) {
		count = 0
	}
	newStr += string(str[0])
	return result, spaceStr, count, newStr
}
