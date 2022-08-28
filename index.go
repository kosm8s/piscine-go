package piscine

func Index(s string, toFind string) int {
	if len(toFind) == 0 || len(s) == 0 {
		return 0
	}
	if len(s) > len(toFind) {
		for i := 0; i < len(s); i++ {
			if s[i] == toFind[0] {
				count := 0
				for j, k := 0, i; j < len(toFind); j, k = j+1, k+1 {
					if toFind[j] == s[k] {
						count++
					}
					if count == len(toFind) {
						return i
					}
				}
			}
		}
	}
	return -1
}
