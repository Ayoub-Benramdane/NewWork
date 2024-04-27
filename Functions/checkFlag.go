package function

func CheckFlag(i int, str, stf string, k, ch, ch1 int) (int, int, int, string) {
	for j := i; j < len(str); j++ {
		ch1++
		stf += string(str[j])
		if str[j] == ')' {
			ch = 1
			break
		}
		if str[j] == ' ' {
			ch = 2
			k = j + 1
			break
		}
	}
	return k, ch, ch1, stf
}
