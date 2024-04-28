package function

func SpacesFlag(result []string, str, newStr, spaceStr string, i, ch, ch1, k int) ([]string, string, string, int) {
	if spaceStr != "" {
		result = append(result, spaceStr)
		spaceStr = ""
		if i+ch1< len(str) && str[i+ch1] == ' ' {
			i = i+ch1+1
		}
	}
	return result, spaceStr, newStr, i
}
