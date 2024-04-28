package function

import (
	"strconv"
)

func ModificationByFlags(result []string, spaceStr, stf string) ([]string, string) {
	IsModified := false
	strToConverted := ""
	strPonc := ""
	for i := len(result) - 1; i >= 0; i-- {
		if IsModified {
			break
		}
		if stf == "(hex)" || stf == "(bin)" {
			for _, c := range result[i] {
				if IsAlphaNumeric(c) {
					IsModified, strToConverted, strPonc = IsHexBin(result[i], stf)
					if IsModified {
						result[i] = strconv.Itoa(int(ConvDic(strToConverted, stf)))
						if strPonc != "" {
							result[i] += strPonc + " "
							strPonc = ""
						}
					}
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
	return result, spaceStr
}
