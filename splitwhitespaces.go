package piscine

func SplitWhiteSpaces(s string) []string {
	var str []string
	var currentWord string

	for _, v := range s {
		if v == ' ' {
			if currentWord != "" {
				str = append(str, currentWord)
				currentWord = ""
			}
		} else {
			currentWord += string(v)
		}
	}

	if currentWord != "" {
		str = append(str, currentWord)
	}

	return str
}
