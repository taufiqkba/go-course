package main

import "fmt"

func main() {

	// constant must first initialize value
	const phi = 3.14
	fmt.Println("Phi in circle value is: ", phi)

	// constant multiple variable
	const (
		firstName string = "Taufiq"
		lastName         = "Kurniawan"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)

}
