package main

import "fmt"

type Pessoa struct {
	Name  string
	Age   int
	Email string
}

func main() {
	// Assigning a struct to a variable
	person := Pessoa{Name: "Matheus", Age: 21, Email: "matheus@gmail.com"}

	fmt.Println("Hi, I am", person.Name)
	fmt.Println("I am,", person.Age, "years old")
}
