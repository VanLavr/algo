package main

import (
	"fmt"
)

func main() {
	var a = []int{1234, 345, 4, 234, 5, 9090, 23, 8721}
	fmt.Println(a)
	b := MergeSort(a)
	fmt.Println(b)
}

func MergeSort(array []int) []int {
	if len(array) < 2 {
		return array
	}
	middle := len(array) / 2
	left := make([]int, middle)
	right := make([]int, middle+len(array)%2)

	iter := 0
	for index, value := range array {
		if index < len(array)/2 {
			left[index] = value
		} else {
			right[iter] = value
			iter++
		}
	}

	left = MergeSort(left)
	right = MergeSort(right)
	result := Merge(array, left, right)
	return result
}

func Merge(origin, l_array, r_array []int) []int {
	i := 0
	j := 0
	iterator := 0
	for i < len(l_array) && j < len(r_array) {
		if l_array[i] < r_array[j] {
			origin[iterator] = l_array[i]
			i++
			iterator++
		} else {
			origin[iterator] = r_array[j]
			j++
			iterator++
		}
	}

	for i < len(l_array) {
		origin[iterator] = l_array[i]
		i++
		iterator++
	}

	for j < len(r_array) {
		origin[iterator] = r_array[j]
		j++
		iterator++
	}

	return origin
}
