package piscine

func ToLower(s string) string {
	result := []rune(s)
	for i := 0; i < len(s); i++ {
		if result[i] >= 'A' && result[i] <= 'Z' {
			result[i] += 32
		}
	}
	return string(result)
}
