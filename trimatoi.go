package piscine

func Atoi(s string) int {
	res := 0
	sign := 1

	for i, v := range s {
		if v == '-' {
			if i == 0 {
				sign = -1
			} else {
				return 0
			}
		} else if v == '+' {
			if i == 0 {
				sign = 1
			} else {
				return 0
			}
		} else if v >= '0' && v <= '9' {
			res = res*10 + int(v) - '0'
		} else {
			return 0
		}
	}
	return res * sign
}

func TrimAtoi(s string) int {
	if len(s) == 0 {
		return 0
	}
	str := ""
	for _, v := range s {
		if v == '-' && len(str) == 0 {
			str += "-"
		}
		if v < '0' || v > '9' {
			continue
		}
		str += string(v)
	}
	if len(str) == 0 || str == "-" {
		return 0
	}
	return Atoi(str)
}
