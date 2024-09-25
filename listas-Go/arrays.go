package main

import "fmt"

func main() {
	// First exemple of declaration array
	var array [2]string
	array[0] = "Hello"
	array[1] = "World"
	// fmt.Println(array)

	// Seccond example of declaration array
	numberPrimos := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(numberPrimos)
	fmt.Println(numberPrimos[1:3])
}
