package piscine

func BasicJoin(elems []string) string {
	var str string
	for _, a := range elems {
		str += a
	}
	return str
}
