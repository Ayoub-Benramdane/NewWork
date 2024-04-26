package function

func NewStr(result []string, spaceStr, newStr, str string, count int) ([]string, string, string) {
	if newStr != "" {
		result = append(result, newStr)
	}
	newStr = ""
	if count == 0 {
		spaceStr += string(str[0])
	}
	return result, newStr, spaceStr
}
