package main

import (
	"fmt"
	"sort"
)

func main() {
	results := SortNumbers([]int{1, 3, 9, 7, 2})
	fmt.Println(results)
}

func SortNumbers(numbers []int) []int {
	if len(numbers) == 0 {
		return []int{}
	}
	sort.Ints(numbers)
	return numbers
}
