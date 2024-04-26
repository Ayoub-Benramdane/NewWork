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
		if result[i+1][0] == ' ' {
			result[i+1] = ""
		}
		if result[i+count-1][0] == ' ' {
			result[i+count-1] = ""
		}
		i += count
	}
	return result, i
}
