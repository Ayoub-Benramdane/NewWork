package function

func ModificationByParamsFlags(result []string, stf string, index int) []string {
	count := 0
	for i := len(result) - 1; i >= 0; i-- {
		if index > 0 {
			if stf == "(cap, " {
				for _, c := range result[i] {
					if IsAlphaNumeric(c) {
						result[i] = Cap(result[i])
						count++
						index--
						break
					}
				}
			}
			if stf == "(up, " {
				for _, c := range result[i] {
					if IsAlphaNumeric(c) {
						result[i] = Up(result[i])
						count++
						index--
						break
					}
				}
			}
			if stf == "(low, " {
				for _, c := range result[i] {
					if IsAlphaNumeric(c) {
						result[i] = Low(result[i])
						count++
						index--
						break
					}
				}
			}
		}
	}
	return result
}
