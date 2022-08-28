package piscine

func BasicAtoi2(s string) int {
	q := []rune(s)
	a := 0

	for b := 0; b < len(s); b++ {
		if q[b] > '9' || q[b] < '0' {
			return 0
		}
		a = a * 10
		a = a + int(q[b]-48)
	}
	return a
}
