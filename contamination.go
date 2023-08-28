package main

import "fmt"

func main() {
	result := Contamination("", "")
	fmt.Println(result)
}

func Contamination(text, char string) string {
	if text == "" || char == "" {
		return ""
	}

	result := ""
	for i := 0; i < len(text); i++ {
		result += char
	}
	return result
}
