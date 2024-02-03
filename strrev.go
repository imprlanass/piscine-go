package piscine

func StrRev(s string) string {
	var revStr string

	for i := 0; i < len(s); i++ {
		revStr = string(s[i]) + revStr
	}

	return revStr
}
