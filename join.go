package piscine

func Join(strs []string, sep string) string {
	var a string
	for i := 0; i < len(strs); i++ {
		if i < len(strs)-1 {
			a += strs[i] + sep
		} else {
			a = a + strs[i]
		}
	}
	return a
}
