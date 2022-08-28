package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	}

	x := 1
	for i := 1; i <= nb; i++ {
		x *= i
		if x < 0 {
			return 0
		}
	}
	return x
}
