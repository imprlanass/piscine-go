package piscine

func Abort(a, b, c, d, e int) int {
	arr := []int{a, b, c, d, e}
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	median := arr[2]
	if len(arr)%2 == 0 {
		median = arr[1] + arr[3]/2.0
	}
	return median
}
