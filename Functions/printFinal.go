package function

func PrintStrFinal(result []string) string {
	for i := 0; i < len(result); i++ {
		if result[i] == "'" {
			result, i = QuotsModifed(result, i)
		}
	}
	res := ""
	res = CheckVowel(result, res)
	return res
}
