package function

import "strconv"

func ModificationByFlags(result []string, stf string) []string {
	IsModified := false
	for i := len(result) - 1; i >= 0; i-- {
		if IsModified {
			break
		}
		if stf == "(hex)" {
			for _, c := range result[i] {
				if IsAlphaNumeric(c) {
					if IsHex(result[i]) {
						result[i] = strconv.Itoa(int(ConvDic(result[i], stf)))
					}
					IsModified = true
					break
				}
			}
		}
		if stf == "(bin)" {
			for _, c := range result[i] {
				if IsAlphaNumeric(c) {
					if IsBin(result[i]) {
						result[i] = strconv.Itoa(int(ConvDic(result[i], stf)))
					}
					IsModified = true
					break
				}
			}
		}
		if stf == "(cap)" {
			for _, c := range result[i] {
				if IsAlphaNumeric(c) {
					result[i] = Cap(result[i])
					IsModified = true
					break
				}
			}
		}
		if stf == "(up)" {
			for _, c := range result[i] {
				if IsAlphaNumeric(c) {
					result[i] = Up(result[i])
					IsModified = true
					break
				}
			}
		}
		if stf == "(low)" {
			for _, c := range result[i] {
				if IsAlphaNumeric(c) {
					result[i] = Low(result[i])
					IsModified = true
					break
				}
			}
		}
	}
	return result
}
