package piscine

func IsPrintable(s string) bool {
	a := []rune(s)
	for i := range a {
		if a[i] >= 32 && a[i] <= 126 {
		} else {
			return false
		}
	}
	return true
}
