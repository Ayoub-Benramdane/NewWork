package function

func QuotsModifed(result []string, i int) ([]string, int) {
	quots := false
	count := 0
	X := result[i-1]
	for j := i + 1; j < len(result); j++ {
		count++
		if result[j] == "'" {
			quots = true
			break
		}
	}
	if quots {
		if len(result[i+1]) == 0 || result[i+1][0] == ' ' {
			result[i+1] = ""
		}
		if i > 0 && len(X) != 0 {
			if len(X) > 2 && !IsPonc(rune(X[len(X)-2])) && X[0] != ' ' && X[0] != '\'' {
				X += " "
			}
		}
		if len(result[i+count-1]) == 0 || result[i+count-1][0] == ' ' {
			result[i+count-1] = ""

		}
		if i < len(result)-count-1 && len(result[i+count+1]) != 0 {
			if !IsPonc(rune(result[i+count+1][0])) && result[i+count+1][0] != ' ' {
				result[i+count] += " "
			}
		}
		for _, c := range result[i+count-1] {
			if IsPonc(c) {
				result[i+count-1] = result[i+count-1][:len(result[i+count-1])-1]
			}
		}
		i += count
	}
	return result, i
}
