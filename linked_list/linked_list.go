package main

import (
	"fmt"
	"sort"
)

var cycle = make(map[*Edge]bool)

type Edge struct {
	Weight int
	Next   *Edge
}

func (e *Edge) PrintLinkedList() {
	nodes := make(map[*Edge]struct{})
	last := new(Edge)
	cycle := false

	for n := e; n.Next != nil; n = n.Next {
		if _, ok := nodes[n]; ok {
			fmt.Printf("[%d] <= cycle\n", n.Weight)
			cycle = true
			break
		} else {
			fmt.Printf("[%d]->", n.Weight)
			nodes[n] = struct{}{}
			last = n.Next
		}
	}

	if !cycle {
		fmt.Printf("[%d]\n", last.Weight)
	}
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
	D := new(Edge)
	E := new(Edge)
	F := new(Edge)
	G := new(Edge)

	A.Weight = 8
	B.Weight = 4
	C.Weight = 1
	D.Weight = 12
	E.Weight = 3
	F.Weight = -2
	G.Weight = 6

	A.Next = B
	B.Next = C
	C.Next = D
	D.Next = E
	E.Next = F
	F.Next = G
	G.Next = nil

	fmt.Println(A.WeightOfTheList())
	A.PrintLinkedList()
	A = SortLinkedListAndDeleteCycle(A)
	A.PrintLinkedList()
}

type SortByValue []Edge

func (a SortByValue) Len() int           { return len(a) }
func (a SortByValue) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortByValue) Less(i, j int) bool { return a[i].Weight < a[j].Weight }

func SortLinkedListAndDeleteCycle(node *Edge) *Edge {
	var Edges SortByValue
	nodes := make(map[*Edge]struct{})
	last := new(Edge)
	cycle := false

	for n := node; n.Next != nil; n = n.Next {
		if _, ok := nodes[n]; ok {
			cycle = true
			break
		} else {
			Edges = append(Edges, *n)
			last = n.Next
			nodes[n] = struct{}{}
		}
	}

	if !cycle {
		Edges = append(Edges, *last)
	}

	sort.Sort(Edges)

	for i := 0; i < Edges.Len()-1; i++ {
		Edges[i].Next = &Edges[i+1]
	}
	Edges[Edges.Len()-1].Next = nil

	return &Edges[0]
}
