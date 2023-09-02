package main

import "fmt"

func main() {
	result := CountRedBeads(4)
	fmt.Println(result)
}

func CountRedBeads(n int) int {
	if n < 2 {
		return 0
	}
	return n*2 - 2
}
