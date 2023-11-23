package numbers

func Sum(nums ...int) int {
	var totalRes int

	for _, num := range nums {
		totalRes += num
	}

	return totalRes
}

func ExtendSliceIfSumEven(s *[]int, valToAppend int) {
	if len(*s)%2 == 0 {
		*s = append(*s, valToAppend)
	}
}
