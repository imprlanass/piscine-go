package piscine

func Compact(ptr *[]string) int {
	count := 0
	idx := 0
	for _, v := range *ptr {
		if v != "" {
			(*ptr)[idx] = v
			idx++
			count++
		}
	}
	*ptr = (*ptr)[:idx]
	return count
}
