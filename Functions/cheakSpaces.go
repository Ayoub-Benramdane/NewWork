package function

func SpacesFlag(result []string, str, newStr, spaceStr string, i, paramOrNot, endFlag, endFlagParam int) ([]string, string, string, int) {
	if spaceStr != "" && paramOrNot == 1 {
		result = append(result, spaceStr)
		spaceStr = ""
		if i+endFlag+1 < len(str) && str[i+endFlag] == ' ' && !IsPonc(rune(str[i+endFlag+1])) {
			i += 1
		}
	} else if spaceStr != "" {
		result = append(result, spaceStr)
		spaceStr = ""
		if endFlagParam+1 < len(str) && str[endFlagParam] == ' ' && !IsPonc(rune(str[endFlagParam+1])) {
			i += 1
		}
	}
	return result, spaceStr, newStr, i
}
