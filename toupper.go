package piscine

func ToUpper(s string) string {
	a := []rune(s)
	res := ""
	for i := range a {
		if a[i] >= 97 && a[i] <= 122 {
			res += string(a[i] - 32)
		} else {
			res += string(a[i])
		}
	}
	return res
}
