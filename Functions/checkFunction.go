package function

func IsPonc(c rune) bool {
	ponc := []rune{',', ';', '.', ':', '!', '?'}
	for _, v := range ponc {
		if c == v {
			return true
		}
	}
	return false
}

func IsAlphaNumeric(str rune) bool {
	if str >= '0' && str <= '9' || str >= 'a' && str <= 'z' || str >= 'A' && str <= 'Z' {
		return true
	}
	return false
}

func IsNum(str rune) bool {
	if str >= '0' && str <= '9' {
		return true
	}
	return false
}

func IsVowel(c rune) bool {
	vowel := []rune{'a', 'e', 'u', 'i', 'o', 'h', 'A', 'E', 'U', 'I', 'O', 'H'}
	for _, v := range vowel {
		if c == v {
			return true
		}
	}
	return false
}
