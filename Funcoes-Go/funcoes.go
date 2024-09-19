package main

import "fmt"

func main() {
	resultado := soma(10, 6)
	fmt.Println("Sua soma é:", resultado)
}

func soma(x int, y int) int {
	return x + y
}
