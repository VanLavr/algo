package main

import (
    "fmt"
)

func main() {
    var a = []int{1, 3, 2, 4, 2, 6, -9, -4, -5}
    fmt.Println(a)
    b := BubbleSort(a)
    fmt.Println(b)
}

func BubbleSort(array []int) []int {
    for i := 0; i < len(array); i++ {
        for j := 0; j < len(array); j++ {
            if array[i] < array[j] {
                array[i], array[j] = array[j], array[i]
            }
        }
    }

    return array
}
