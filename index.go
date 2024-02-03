package piscine

func Index(s string, toFind string) int {
	if len(toFind) == 0 {
		return 0
	}
	if len(s) == 0 {
		return -1
	}
	for iS := range s {
		if s[iS] == toFind[0] {
			if len(toFind) == 1 {
				return iS
			}
			if len(s[iS:]) >= len(toFind) {
				if s[iS:iS+len(toFind)] == toFind {
					return iS
				}
			}
		}
		iS++
	}
	return -1
}
