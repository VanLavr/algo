package main

import (
	"fmt"
)

type MySet[E comparable] map[E]struct{}

type Human struct {
	Name string
}

func NewSet[E comparable](args ...E) MySet[E] {
	newSet := MySet[E]{}
	for _, arg := range args {
		newSet[arg] = struct{}{}
	}
	return newSet
}

func (s MySet[E]) Add(values ...E) {
	for _, value := range values {
		s[value] = struct{}{}
	}
}

func (s MySet[E]) Contains(value E) bool {
	_, ok := s[value]
	return ok
}

func (s MySet[E]) Members() []E {
	var result []E
	for member, _ := range s {
		result = append(result, member)
	}
	return result
}

func main() {
	testWithints := NewSet(1, 3, 5)
	testWithbools := NewSet(true, false, false)
	testWithstructs := NewSet(Human{Name: "Joe"}, Human{Name: "Ivan"})

	testWithints.Add(0)
	testWithbools.Add(true)
	testWithstructs.Add(Human{Name: "Vasya"})
	fmt.Println(testWithints.Contains(1))
	fmt.Println(testWithstructs, testWithbools, testWithints)
}
