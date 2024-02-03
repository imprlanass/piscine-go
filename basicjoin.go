package piscine

func BasicJoin(elems []string) string {
	str := ""
	for _, v := range elems {
		str += string(v)
	}
	return str
}
