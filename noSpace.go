package main

import (
	"fmt"
	"strings"
)

func NoSpace(word string) string {
	return strings.ReplaceAll(word, " ", "")
}

func main() {
	inputString := "Пример строки с пробелами"
	stringWithoutSpaces := NoSpace(inputString)
	fmt.Println(stringWithoutSpaces)
}
