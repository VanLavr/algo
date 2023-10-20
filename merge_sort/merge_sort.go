package main

import (
    "fmt"
)

func main() {
    var a = []int{1, 3, 2, 5, 3, 4, 345, 234, 5, 23}
    b := MergeSort(a)
    fmt.Println(b)
}

func MergeSort(array []int) []int {
    middle := len(array) / 2
    left := make([]int, middle)
    right := make([]int, middle + len(array) % 2)

    for index, value := range array {
        if index < len(array) / 2 {
            left[index] = value
        } else {
            right[index] = value
        }
    }

    MergeSort(left)
    MergeSort(right)
    result := Merge(array, left, right)
    return result
}

func Merge(origin, l_array, r_array []int) []int {
    merged := make([]int, len(l_array) + len(r_array))

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
    }

    for j < len(r_array) {
        origin[iterator] = r_array[j]
    }
    
    return origin
}







