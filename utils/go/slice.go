package utils

func SliceMax(slice []int) (max int) {
	for _, num := range slice {
		if max < num {
			max = num
		}
	}
	return
}

func SliceMin(slice []int) int {
	min := slice[0]
	for _, num := range slice {
		if num < min {
			min = num
		}
	}
	return min
}
