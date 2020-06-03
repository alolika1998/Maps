package main

import "fmt"

func main() {
	statePopulation := map[string]int {
		"California" : 39250017,
		"Texas" :      27862596,
		"New York" :   19745289,
	}
	fmt.Println(statePopulation)
	fmt.Println(statePopulation["Texas"])
	statePopulation["florida"] = 11612432
	fmt.Println(statePopulation)
	fmt.Println(len(statePopulation))
	delete(statePopulation, "florida")
	fmt.Println(statePopulation)
}
