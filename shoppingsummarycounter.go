package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	iMap := make(map[string]int)

	wc := SplitWhiteSpaces(str)

	for _, v := range wc {
		iMap[v]++
	}
	return iMap
}
