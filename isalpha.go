package piscine

func IsAlpha(s string) bool {
	a := []rune(s)
	for i := range a {
		if a[i] >= 97 && a[i] <= 122 || a[i] >= 65 && a[i] <= 90 || a[i] >= 48 && a[i] <= 57 {
		} else {
			return false
		}
	}
	return true
}
