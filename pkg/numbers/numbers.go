package numbers

func Sum(nums ...int) int {
	var totalRes int

	for _, num := range nums {
		totalRes += num
	}

	return totalRes
}

//2 Task
//	Does it work?
//	No, it's not. Because in the first version of signature we're passing
//	copy of array and appending value to the COPY.
//	Are there an extra condition that slice should fulfill?
//	I don't think so, we do not need to worry about memory allocation because append
//	handles cases when cap > len.
//	How the function signature should change?
//	We should return modified copy of slice or use a pointer.

func ExtendSliceIfLenEven(s *[]int, valToAppend int) {
	if len(*s)%2 == 0 {
		*s = append(*s, valToAppend)
	}
}
