package main

import "fmt"

func main() {
	// Maps using key and value
	ages := make(map[string]int)

	ages["Marcos"] = 25
	ages["Daniel"] = 29
	ages["Milena"] = 19
	ages["Michael"] = 21

	fmt.Println(ages)

	// Show ages values by the key
	fmt.Println(ages["Milena"])
}
