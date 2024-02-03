package piscine

func ToUpper(s string) string {
	// var v rune
	// for _, v = range s {
	// 	if 97 <= v && v <= 122 {
	// 		v = v - 32
	// 	}
	// }
	// return string(v)
	result := []rune(s)
	for i := 0; i < len(s); i++ {
		if result[i] >= 'a' && result[i] <= 'z' {
			result[i] -= 32
		}
	}
	return string(result)
}
