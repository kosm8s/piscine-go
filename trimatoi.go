package piscine

func TrimAtoi(s string) int {
	q := []rune(s)
	a := 0
	firstdigit := false
	minus := false
	if s == "" {
		return 0
	}

	for b := 0; b < len(s); b++ {

		if q[b] == '-' && !firstdigit {
			minus = true
		}
		if q[b] <= '9' && q[b] >= '0' {
			firstdigit = true
		}
		if b == 0 && minus {
			continue
		}
		if q[b] > '9' || q[b] < '0' {
			continue
		}

		a = a * 10
		a = a + int(q[b]-48)

	}
	if minus {
		a = -a
	}
	return a
}
