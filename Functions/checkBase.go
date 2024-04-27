package function

func IsHexBin(str, base string) (bool, string, string) {
	var baseToConverte []rune
	strToConverted := ""
	strPonc := ""
	breaking := false
	if base == "(bin)" {
		baseToConverte = []rune{'0', '1'}
	} else {
		baseToConverte = []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f', 'A', 'B', 'C', 'D', 'E', 'F'}
	}
	IsValid := false
	for i, c := range str {
		if breaking {
			break
		}
		IsValid = false
		for _, v := range baseToConverte {
			if strPonc != "" && c == ' ' && i == len(str)-1 {
				breaking = true
				IsValid = true
				break
			}
			if c == v || IsPonc(c) {
				if !IsPonc(c) {
					strToConverted += string(c)
				} else if IsPonc(c) {
					strPonc += string(c)
				}
				IsValid = true
				break
			}
		}
	}
	if IsValid {
		return IsValid, strToConverted, strPonc
	}
	return false, "", ""
}
