package main

import "fmt"

func main() {
	result := Between(-1, 8)
	fmt.Println(result)
}

func Between(a, b int) []int {
	var nums []int
	for i := a; i <= b; i++ {
		nums = append(nums, i)
	}
	return nums
}
