package piscine

func IsUpper(s string) bool {
	a := []rune(s)
	for i := range a {
		if a[i] >= 65 && a[i] <= 90 {
		} else {
			return false
		}
	}
	return true
}
