package main

import "fmt"

func add(n int) func(int) int {
	return func(x int) int {
		return x + n
	}
}

func main() {
	add := add(5)
	result := add(3)
	fmt.Println(result)
}
