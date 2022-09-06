package main

import "fmt"

func main() {

	// variable with data type
	var name string
	name = "Taufiq Kurniawan"
	fmt.Println(name)

	name = "Bayu Amartha"
	fmt.Println(name)

	// variable without data type
	var friendName = "Taufiq Kurniawan \n"
	fmt.Println(friendName)

	var age = 21
	fmt.Println(age)

	// variable without "var"
	country := "Indonesia"
	fmt.Println(country)

	country = "United States"
	fmt.Println(country)

	// multiple variable
	var (
		firstName = "Koko"
		lastName  = "Kuliner"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)

}
