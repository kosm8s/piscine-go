package piscine

func Unmatch(a []int) int {
	for i, ind := range a {
		count := 1
		for j, jn := range a {
			if ind == jn && i != j {
				count++
			}
		}
		if count%2 == 1 {
			return ind
		}
	}
	return -1
}
