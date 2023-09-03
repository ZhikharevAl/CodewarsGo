package main

import "fmt"

func main() {
	result := SumCubes(10)
	fmt.Println(result)
}

func SumCubes(n int) int {
	sum := 0
	for i := 0; i <= n; i++ {
		sum += i * i * i
	}
	return sum
}
