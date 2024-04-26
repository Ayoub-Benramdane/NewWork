package function

func IsHex(str string) bool {
	base := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f', 'A', 'B', 'C', 'D', 'E', 'F'}
	IsValid := false
	for _, c := range str {
		IsValid = false
		for _, v := range base {
			if c == v {
				IsValid = true
				break
			}
		}
	}
	if IsValid {
		return IsValid
	}
	return false	
}

func IsBin(str string) bool {
	base := []rune{'0', '1'}
	IsValid := false
	for _, c := range str {
		IsValid = false
		for _, v := range base {
			if c == v {
				IsValid = true
				break
			}
		}
	}
	if IsValid {
		return IsValid
	}
	return false	
}