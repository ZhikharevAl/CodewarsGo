package main

import "fmt"

func main() {
	x := "45385593107843568"
	result := FakeBin(x)
	fmt.Println(result)
}

func FakeBin(x string) string {
	b := []byte(x)
	for i := range b {
		if b[i] < '5' {
			b[i] = '0'
		} else {
			b[i] = '1'
		}
	}
	return string(b)
}
