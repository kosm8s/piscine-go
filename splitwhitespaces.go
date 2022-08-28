package piscine

func SplitWhiteSpaces(s string) []string {
	var ans []string
	q := ""
	for i := 0; i < len(s); i++ {
		q = ""
		for s[i] != ' ' && s[i] != '\n' && s[i] != 9 {
			q += string(s[i])
			if i == len(s)-1 {
				break
			}
			i++
		}
		if q != "" {
			ans = append(ans, q)
		}

	}
	return ans
}
