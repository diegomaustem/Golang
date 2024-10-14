package main

import "fmt"

func main() {

	sum := 0

	// Normal for example
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("Fora do loop:", sum)

	// For range example
	frutas := []string{"laranja", "maça", "banana", "uva", "Kiwi"}
	for _, fruta := range frutas {
		fmt.Println(fruta)
	}
}
