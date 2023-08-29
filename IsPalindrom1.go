package main

import (
	"fmt"
	"strings"
)

func main() {
	result := IsPal("Abba")
	fmt.Println(result)
}

func IsPal(str string) bool {
	reversed := ""
	str = strings.ToLower(str)
	for i := len(str) - 1; i >= 0; i-- {
		reversed += string(str[i])
	}
	return str == reversed
}
