package function

func Split(str string) []string {
	result := []string{}
	spaceStr := ""
	newStr := ""
	IsValid := false
	count := 0
	for i := 0; i < len(str); i++ {
		if count == 0 {
			if IsPonc(rune(str[i])) {
				count++
				spaceStr = ""
			}
		}
		if IsPonc(rune(str[i])) {
			result, newStr, i = SplitPonc(result, newStr, str, i)
			if i >= len(str) {
				break
			}
		}
		if str[i] == '\'' {
			result, newStr, spaceStr, IsValid = SingleQuots(result, spaceStr, newStr, str, i)
			if IsValid {
				continue
			}
		}
		if str[i] == '(' {
			result, i, newStr, spaceStr = SplitFlags(str, result, i, newStr, spaceStr)
			if i >= len(str) {
				break
			}
		}
		if str[i] != ' ' && str[i] != '\n' {
			result, spaceStr, count, newStr = SpaceStr(result, newStr, spaceStr, str[i:], count)
		}
		if str[i] == '\n' {
			newStr += string(str[i])
			result = append(result, newStr)
			newStr = ""
		}
		if str[i] == ' ' {
			result, newStr, spaceStr = NewStr(result, spaceStr, newStr, str[i:], count)
		}
	}
	if newStr != "" {
		result = append(result, newStr)
	}
	return result
}
