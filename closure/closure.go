package main

import "fmt"

func main() {
	f := Closure(3)
	fmt.Println(f() == "yes")
	fmt.Println(f() == "yes")
	fmt.Println(f() == "yes")
	fmt.Println(f() == "yes")
	fmt.Println(f() == "yes")
}

func Closure(n int) func() string {
	var counter int

	if n < 1 {
		return func() string {
			return "no"
		}
	}

	return func() string {
		counter++
		if counter > n {
			return "no"
		} else {
			return "yes"
		}
	}
}
