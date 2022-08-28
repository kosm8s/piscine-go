package piscine

func Abort(a, b, c, d, e int) int {
	sorted := []int{a, b, c, d, e}
	bubble_sort(sorted)
	return sorted[2]
}

func bubble_sort(table []int) {
	for i := 0; i < len(table); i++ {
		for j := 1; j < len(table); j++ {
			if table[j] < table[j-1] {
				table[j], table[j-1] = table[j-1], table[j]
			}
		}
	}
}
