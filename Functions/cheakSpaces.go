package function

func SpacesFlag(result []string, str, newStr, spaceStr string, i, ch, ch1, k int) ([]string, string, string, int) {
	if spaceStr != "" && ch == 1 {
		result = append(result, spaceStr)
		spaceStr = ""
		if i+ch1+1 < len(str) && str[i+ch1] == ' ' && !IsPonc(rune(str[i+ch1+1])) {
			i += 1
		}
	} else if spaceStr != "" {
		result = append(result, spaceStr)
		spaceStr = ""
		if k+1 < len(str) && str[k] == ' ' && !IsPonc(rune(str[k+1])) {
			i += 1
		}
	}
	return result, spaceStr, newStr, i
}
