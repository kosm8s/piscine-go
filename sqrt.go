package piscine

func Sqrt(nb int) int {
	if nb < 0 {
		return 0
	} else {
		for i := 0; i < 999; i++ {
			if nb == i*i {
				nb = i
				break
			} else if nb < i*i {
				return 0
			}
		}
	}
	return nb
}
