package function

func SingleQuots(result []string, spaceStr, newStr, str string, i int) ([]string, string, string, bool) {
	IsValid := true
	for j := i; j < len(str); j++ {
		if i != 0 && i != len(str)-1 {
			if str[i-1] != ' ' && str[i-1] != '\n' && str[i+1] != ' ' && str[i+1] != '\n' {
				IsValid = false
				break
			}
		} else {
			IsValid = true
			break
		}
	}
	if IsValid {
		if newStr != "" || spaceStr != "" {
			result = append(result, newStr, spaceStr)
			newStr = ""
			spaceStr = ""
		}
		result = append(result, string(str[i]))
	}
	return result, newStr, spaceStr, IsValid
}
