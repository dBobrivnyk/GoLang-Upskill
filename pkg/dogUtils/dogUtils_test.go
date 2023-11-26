package dogUtils_test

import (
	"github.com/dbobrivnykObj/GoLang-Upskill/pkg/dogUtils"
	"testing"
)

var aDog = dogUtils.Dog{Name: "Alabaster", Age: 12}
var bDog = dogUtils.Dog{Name: "Bulldog", Age: 5}
var cDog = dogUtils.Dog{Name: "Mastiff", Age: 8}

func TestSortDogByName(t *testing.T) {
	//Arrange
	dogs := []dogUtils.Dog{aDog, cDog, bDog}
	expected := []dogUtils.Dog{aDog, bDog, cDog}
	//Act
	dogUtils.SortDogByName(dogs)
	//Assert
	for i := range expected {
		if expected[i] != dogs[i] {
			t.Errorf("Dog == %q, want %q", dogs[i], expected[i])
		}
	}
}

func TestSortDogPtrByName(t *testing.T) {
	//Arrange
	dogs := []*dogUtils.Dog{&aDog, &cDog, &bDog}
	expected := []*dogUtils.Dog{&aDog, &bDog, &cDog}
	//Act
	dogUtils.SortDogPtrByName(dogs)
	//Assert
	for i := range expected {
		if expected[i] != dogs[i] {
			t.Errorf("Dog == %q, want %q", dogs[i], expected[i])
		}
	}
}

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
