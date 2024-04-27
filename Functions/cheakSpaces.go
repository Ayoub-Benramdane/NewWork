package function

func SpacesFlag(str, newStr, spaceStr string, i, ch, ch1, k int) (string, string, int) {
	if len(spaceStr) > 1 {
		spaceStr = spaceStr[0 : len(spaceStr)-1]
	} else {
		spaceStr = ""
	}
	if ch == 1 && str[i+ch1] == ' ' {
		i += 1
	}
	if ch == 2 && str[k+1] != ' ' {
		newStr += " "
	}
	return spaceStr, newStr, i
}
