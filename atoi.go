package piscine

func Atoi(s string) int {
	q := []rune(s)
	a := 0
	minus := false
	plus := false
	if s == "" {
		return 0
	}

	if q[0] == '-' {
		minus = true
	}
	if q[0] == '+' {
		plus = true
	}

	for b := 0; b < len(s); b++ {
		if b == 0 && minus {
			continue
		}
		if b == 0 && plus {
			continue
		}
		if q[b] > '9' || q[b] < '0' {
			return 0
		}
		a = a * 10
		a = a + int(q[b]-48)
	}
	if minus {
		a = -a
	}
	return a
}
