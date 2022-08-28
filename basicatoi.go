package piscine

func BasicAtoi(s string) int {
	q := []rune(s)
	a := 0

	for b := 0; b < len(s); b++ {
		a = a * 10
		a = a + int(q[b]-48)
	}
	return a
}
