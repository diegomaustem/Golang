package main

import "fmt"

func main() {
	numberOne := 2
	numberTwo := 3
	operation := "sum"

	resultOperation := calculator(numberOne, numberTwo, operation)
	if resultOperation == -1 {
		fmt.Println("Error in operation")
	} else {
		fmt.Println("Your", operation, "is", resultOperation)
	}
}

func calculator(numberOne int, numberTwo int, operation string) int {
	if operation == "sum" {
		return numberOne + numberTwo
	} else if operation == "sub" {
		return numberOne - numberTwo
	} else if operation == "mult" {
		return numberOne * numberTwo
	} else if operation == "div" {
		return numberOne / numberTwo
	} else {
		return -1
	}
}
