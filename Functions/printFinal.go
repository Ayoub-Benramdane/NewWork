package function

func PrintStrFinal(result []string) string {
	for i := 0; i < len(result); i++ {
		if result[i] == "'" {
			result, i = QuotsModifed(result, i)
		}
	}
	res := ""
	for i := 0; i < len(result); i++ {
		if i < len(result)-1 && (result[i] == "a" || result[i] == "A") {
			l := i + 1
			for j := l; j < len(result); j++ {
				if result[j] != "" && result[j][0] != ' ' {
					if IsVowel(rune(result[j][0])) {
						res += "an"
					} else {
						res += result[i]
					}
					break
				}
			}
		} else if i < len(result)-1 && (result[i] == "an" || result[i] == "An") {
			l := i + 1
			for j := l; j < len(result); j++ {
				if result[j] != "" && result[j][0] != ' ' {
					if !IsVowel(rune(result[j][0])) {
						res += "a"
					} else {
						res += result[i]
					}
					break
				}
			}
		} else {
			res += result[i]
		}
	}
	return res
}
