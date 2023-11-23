package dogUtils_test

import (
	"github.com/dbobrivnykObj/GoLang-Upskill/pkg/dogUtils"
	"testing"
)

//Tested by IDE - PTR is quicker
/*
BenchmarkSortDogByName
BenchmarkSortDogByName-12                4884068               248.6 ns/op
BenchmarkSortDogPtrByName
BenchmarkSortDogPtrByName-12             5599435               205.2 ns/op
PASS
*/

var aDog = dogUtils.Dog{Name: "Husky", Age: 12}
var bDog = dogUtils.Dog{Name: "Bulldog", Age: 5}
var cDog = dogUtils.Dog{Name: "Mastiff", Age: 8}

func BenchmarkSortDogByName(b *testing.B) {
	dogs := []dogUtils.Dog{aDog, bDog, cDog}

	for i := 0; i < b.N; i++ {
		dogUtils.SortDogByName(dogs)
	}
}

func BenchmarkSortDogPtrByName(b *testing.B) {
	dogsPtr := []*dogUtils.Dog{&aDog, &bDog, &cDog}

	for i := 0; i < b.N; i++ {
		dogUtils.SortDogPtrByName(dogsPtr)
	}
}
