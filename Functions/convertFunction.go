package function

import "strconv"

func ConvDic(str, base string) int64 {
	if base == "(hex)" {
		res, _ := strconv.ParseInt(str, 16, 64)
		return res
	} else if base == "(bin)" {
		res, _ := strconv.ParseInt(str, 2, 64)
		return res
	}
	return 0
}

func Up(str string) string {
	res := ""
	for _, c := range str {
		if c >= 'a' && c <= 'z' {
			c -= 32
		}
		res += string(c)
	}
	return res
}

func Low(str string) string {
	res := ""
	for _, c := range str {
		if c >= 'A' && c <= 'Z' {
			c += 32
		}
		res += string(c)
	}
	return res
}

func Cap(str string) string {
	res := ""
	for i, c := range str {
		if c >= 'a' && c <= 'z' && i == 0 {
			c -= 32
		} else if c >= 'A' && c <= 'Z' {
			c += 32
		}
		res += string(c)
	}
	return res
}
