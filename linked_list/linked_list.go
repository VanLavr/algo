package main

import (
	"fmt"
)

var cycle = make(map[*Edge]bool)

type Edge struct {
	Weight int
	Next   *Edge
}

func (e *Edge) WeightOfTheList() int {
	if _, ok := cycle[e]; ok {
		return 0
	} else {
		cycle[e] = true
	}

	if e.Next == nil {
		return e.Weight
	}

	return e.Weight + e.Next.WeightOfTheList()
}

func main() {
	A := new(Edge)
	B := new(Edge)
	C := new(Edge)

	A.Weight = 8
	B.Weight = 4
	C.Weight = 1

	A.Next = B
	B.Next = C
	C.Next = A

	fmt.Println(A.WeightOfTheList())
}
