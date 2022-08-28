package piscine

func StrLen(s string) int {
	chars := []rune(s)
	counter := 0
	for i := range chars {
		counter = i
	}
	counter++
	return counter
}
