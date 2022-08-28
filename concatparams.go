package piscine

func ConcatParams(args []string) string {
	n := 0
	res := ""
	for range args {
		n++
	}
	for i := 0; i < n-1; i++ {
		res = res + args[i] + "\n"
	}
	res = res + args[n-1]
	return res
}
