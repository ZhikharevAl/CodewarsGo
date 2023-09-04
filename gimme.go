package main

import (
	"fmt"
	"sort"
)

func main() {
	result := Gimme([3]int{5, 3, 8})
	fmt.Println(result)
}

func Gimme(array [3]int) int {

	sortedArray := make([]int, 3)
	copy(sortedArray, array[:])
	sort.Ints(sortedArray)

	middle := sortedArray[1]

	for i, v := range array {
		if v == middle {
			return i
		}
	}

	return -1
}
