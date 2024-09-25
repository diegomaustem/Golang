package main

import "fmt"

func main() {
	// Using slices without fixing number of positions
	var sliceWithoutPositions []string

	// Adding value in slice bellow
	sliceWithoutPositions = append(sliceWithoutPositions, "Employer 01")

	fmt.Println(sliceWithoutPositions)

	// Using slices with fixing number of positions
	sliceWithPositions := make([]string, 5)

	sliceWithPositions[0] = "Hello"
	sliceWithPositions[1] = "World"
	fmt.Println(sliceWithPositions[0], sliceWithPositions[1])

	// Other types of created slices
	numberPair := []int{2, 4, 6, 8, 10}

	// Adding new value in slice bellow
	numberPair = append(numberPair, 12, 14, 16)
	fmt.Println(numberPair)
}
