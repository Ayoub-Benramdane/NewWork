package function

func CheckFlag(i int, str, stf string, endFlagParam, paramOrNot, endFlag int) (int, int, int, string) {
	for j := i; j < len(str); j++ {
		endFlag++
		stf += string(str[j])
		if str[j] == ')' {
			paramOrNot = 1
			break
		}
		if str[j] == ' ' && j != len(str)-1 {
			paramOrNot = 2
			endFlagParam = j + 1
			break
		}
	}
	return endFlagParam, paramOrNot, endFlag, stf
}
