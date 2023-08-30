package main

import (
	"fmt"
)

func main() {
	number := 5
	table := MultiTable(number)
	fmt.Println(table)
}

func MultiTable(number int) string {
	if number < 1 || number > 10 {
		return "Number should be between 1 and 10"
	}

	table := ""
	for i := 1; i <= 10; i++ {
		result := number * i
		table += fmt.Sprintf("%d * %d = %d\n", i, number, result)
	}

	return table
}
