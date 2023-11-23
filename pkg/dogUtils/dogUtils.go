package dogUtils

import (
	"sort"
	"strings"
)

type Dog struct {
	Name string
	Age  uint32
}

func MakeBirthday(dogs []Dog) {
	for i := range dogs {
		dogs[i].Age++
	}
}

func SortDogByName(dogs []Dog) {
	sort.Slice(dogs, func(i, j int) bool {
		return strings.ToLower(dogs[i].Name) < strings.ToLower(dogs[j].Name)
	})
}

func SortDogPtrByName(dogs []*Dog) {
	sort.Slice(dogs, func(i, j int) bool {
		return strings.ToLower(dogs[i].Name) < strings.ToLower(dogs[j].Name)
	})
}
