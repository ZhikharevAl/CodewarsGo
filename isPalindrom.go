package main

import "fmt"

func main() {
	result := IsPalindrom("aba")
	fmt.Println(result)
}

func IsPalindrom(str string) bool {
	length := len(str)
	for i := 0; i < length/2; i++ {
		if str[i] != str[length-1-i] {
			return false
		}
	}
	return true
}
