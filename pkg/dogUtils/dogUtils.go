package dogUtils

import (
	"sort"
	"strings"
	"sync"
)

type Dog struct {
	Name string
	Age  uint32
}

func (a Dog) Compare(b Dog) int {
	if a.Name > b.Name {
		return 1
	}
	if a.Name < b.Name {
		return -1
	}
	if a.Age > b.Age {
		return 1
	}
	if a.Age < b.Age {
		return -1
	}
	return 0
}

func MakeBirthday(dogs []Dog) {
	wg := sync.WaitGroup{}
	for i := range dogs {
		wg.Add(1)
		go func(j int) {
			dogs[j].Age++
			wg.Done()
		}(i)
	}
	wg.Wait()
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

// missing sort interface implementation
// implement sort interface for []Dog structure to sort by age + write ut
// you can define new type for that
// type Dogs []Dog
// see: https://medium.com/@kdnotes/sort-sort-interface-in-golang-1d263d96956d
