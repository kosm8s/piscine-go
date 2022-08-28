package piscine

func IsLower(s string) bool {
	a := []rune(s)
	for i := range a {
		if a[i] >= 97 && a[i] <= 122 {
		} else {
			return false
		}
	}
	return true
}
