package main

import (
	"4d63.com/strrev"
	"fmt"
)

func main() {
	result := IsPalindrome("a")
	fmt.Println(result)
}

func IsPalindrome(str string) bool {
	return str == strrev.Reverse(str)
}
