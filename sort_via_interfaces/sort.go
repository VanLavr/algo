package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type MyInt []int

type Human struct {
	Name string
	Age  int
}

type Crowd []Human

func (a Crowd) Len() int           { return len(a) }
func (a Crowd) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Crowd) Less(i, j int) bool { return a[i].Age < a[j].Age }

func (a MyInt) Len() int           { return len(a) }
func (a MyInt) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a MyInt) Less(i, j int) bool { return a[i] < a[j] }

func main() {
	var a = MyInt{1, 3, 2, 5, 3, 6}
	fmt.Println(a)
	sort.Slice(a, func(i, j int) bool {
		return a[i] > a[j]
	})
	fmt.Println(a)

	// structures

	var c Crowd

	for i := 0; i < 15; i++ {
		var h = Human{
			Name: fmt.Sprintf("Man #%d", i),
			Age:  rand.Intn(50),
		}

		c = append(c, h)
	}

	for _, h := range c {
		fmt.Println(h)
	}

	sort.Slice(c, func(i, j int) bool {
		return c[i].Age < c[j].Age
	})

	fmt.Println("--------")
	for _, h := range c {
		fmt.Println(h)
	}
}
