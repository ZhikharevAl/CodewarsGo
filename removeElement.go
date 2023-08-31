package main

import "fmt"

func main() {
	nums := []int{3, 2, 2, 3}
	val := 3
	result := removeElement(nums, val)
	fmt.Println(result)
}

func removeElement(nums []int, val int) int {
	k := 0

	for _, num := range nums {
		if num != val {
			nums[k] = num
			k++
		}
	}

	return k
}
