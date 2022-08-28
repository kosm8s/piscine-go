package piscine

func IsNumeric(s string) bool {
	a := []rune(s)
	for i := range a {
		if a[i] >= 48 && a[i] <= 57 {
		} else {
			return false
		}
	}
	return true
}
