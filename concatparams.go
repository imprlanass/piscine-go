package piscine

func ConcatParams(args []string) string {
	var s string
	for i := range args {
		s += args[i]
		if i != len(args)-1 {
			s += "\n"
		}
	}
	return s
}
