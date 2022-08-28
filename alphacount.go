package piscine

func AlphaCount(s string) int {
	q := []rune(s)
	count := 0
	for i := range q {
		if q[i] >= 65 && q[i] <= 90 || q[i] >= 97 && q[i] <= 122 {
			count++
		}
	}
	return count
}
