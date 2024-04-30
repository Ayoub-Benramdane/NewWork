package function

func QuotsModifed(result []string, i int) ([]string, int) {
	quots := false
	count := 0
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
		if i > 0 && len(result[i-1]) != 0 {
			if len(result[i-1]) > 2 && !IsPonc(rune(result[i-1][len(result[i-1])-2])) && result[i-1][0] != ' ' && result[i-1][0] != '\'' {
				result[i-1] += " "
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
				break
			}
		}
		i += count
	}
	return result, i
}
