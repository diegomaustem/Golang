package main

import "fmt"

type Person struct {
	Name  string
	Age   int
	Email string
}

// Defined struct with type Person
type Occupation struct {
	Person
	Type string
}

func main() {
	// Assigning a struct to a variable
	person := Person{Name: "Matheus", Age: 21, Email: "matheus@gmail.com"}

	fmt.Println("Hi, I am", person.Name)
	fmt.Println("I am,", person.Age, "years old")

	// Getting struct with type Person
	occupation := Occupation{person, "Developer"}

	fmt.Println("Hi, I am a", occupation.Type)
	fmt.Println("Hi, my name is", occupation.Person.Name)
	fmt.Println("My occupation is", occupation.Person.Age)
}
