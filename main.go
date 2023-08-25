package main

import (
	"fmt"
)

func main() {
	totalGoals := Goals(43, 10, 5)
	fmt.Println(totalGoals)
}

func Goals(laLigaGoals, copaDelReyGoals, championsLeagueGoals int) int {
	return laLigaGoals + copaDelReyGoals + championsLeagueGoals
}
