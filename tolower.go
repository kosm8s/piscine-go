package piscine

func ToLower(s string) string {
	a := []rune(s)
	res := ""
	for i := range a {
		if a[i] >= 65 && a[i] <= 90 {
			res += string(a[i] + 32)
		} else {
			res += string(a[i])
		}
	}
	return res
}
