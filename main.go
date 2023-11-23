package main

import (
	"fmt"
	"github.com/dbobrivnykObj/GoLang-Upskill/pkg/dogUtils"
	"github.com/dbobrivnykObj/GoLang-Upskill/pkg/numbers"
	"github.com/dbobrivnykObj/GoLang-Upskill/pkg/wordUtils"
)

func main() {
	nums := []int{5, 5, 5, 5}
	//1 Task
	fmt.Println("Task 1:", numbers.Sum(1, 2, 3), numbers.Sum(nums...))

	//2 Task
	/*
		Does it work?
		No, it's not. Because in the first version of signature we're passing
		copy of array and appending value to the COPY.
		Are there an extra condition that slice should fulfill?
		I don't think so, we do not need to worry about memory allocation because append
		handles cases when cap > len.
		How the function signature should change?
		We should return modified copy of slice or use a pointer.
	*/
	numbers.ExtendSliceIfSumEven(&nums, 4)
	fmt.Println(nums)

	//3 Task
	a := dogUtils.Dog{"Husky", 12}
	b := dogUtils.Dog{"Bulldog", 5}
	c := dogUtils.Dog{"Mastiff", 8}

	dogs := []dogUtils.Dog{a, b, c}
	dogsPtr := []*dogUtils.Dog{&a, &b, &c}

	dogUtils.MakeBirthday(dogs)
	fmt.Println(dogs)

	dogUtils.SortDogByName(dogs)
	dogUtils.SortDogPtrByName(dogsPtr)
	fmt.Println(dogs)

	for _, v := range dogsPtr {
		fmt.Println(v.Name)
	}

	//4 Task
	fmt.Println(wordUtils.AreAnagram("listen", "silent"))
	fmt.Println(wordUtils.AreAnagram("hand", "hans"))

	//5 Task
	fmt.Println(wordUtils.IsPalindrome("kążąk"))
	fmt.Println(wordUtils.IsPalindrome("samolot"))

}
