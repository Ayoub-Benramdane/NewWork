package function

func CheckVowel(result []string, res string) string {
	for i := 0; i < len(result); i++ {
		if i < len(result)-1 && result[i] == "a" {
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
		} else if i < len(result)-1 && result[i] == "A" {
			l := i + 1
			for j := l; j < len(result); j++ {
				if result[j] != "" && result[j][0] != ' ' {
					if IsVowel(rune(result[j][0])) {
						res += "An"
					} else {
						res += result[i]
					}
					break
				}
			}
		} else if i < len(result)-1 && result[i] == "an" {
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
		} else if i < len(result)-1 && result[i] == "An" {
			l := i + 1
			for j := l; j < len(result); j++ {
				if result[j] != "" && result[j][0] != ' ' {
					if !IsVowel(rune(result[j][0])) {
						res += "A"
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
