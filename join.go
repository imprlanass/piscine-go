package piscine

func Join(strs []string, sep string) string {
	str := ""
	for i := range strs {
		if i != len(strs)-1 {
			str += strs[i] + sep
		} else {
			str += strs[i]
		}
	}
	return str
}
