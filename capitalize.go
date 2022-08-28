package piscine

func Capitalize(s string) string {
	q := []rune(s)
	len := 0
	for range q {
		len++
	}
	for i, b := range q {
		if i == 0 || !isalpha(q[i-1]) {
			if b >= 97 && b <= 122 {
				q[i] = b - 32
			}
		} else {
			if b >= 65 && b <= 90 {
				q[i] = b + 32
			}
		}
	}
	return string(q)
}

func isalpha(a rune) bool {
	if a >= 97 && a <= 122 || a >= 65 && a <= 90 || a >= 48 && a <= 57 {
		return true
	}
	return false
}
