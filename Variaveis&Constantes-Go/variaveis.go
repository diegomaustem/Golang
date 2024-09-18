package main

import "fmt"

func main() {
	var name, lastname string = "Thiago", "Cardoso"
	fmt.Println("Your name is:", name)
	fmt.Println("Your lastname is:", lastname)

	var age int
	age = 22
	fmt.Println("Your age is:", age)

	var height float64
	height = 1.69
	fmt.Println("Your height is:", height)

	// Another type of declaration for variables
	ocupation := "Enginier"
	fmt.Println("He is an:", ocupation)

	salary := 8.600
	fmt.Println("Your salary is:", salary)
}
