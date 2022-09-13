package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Looping into-", counter)
		counter++
	}

	//	for with statement
	for value := 1; value <= 10; value++ {
		fmt.Println("Value to: ", value)
	}

	//	for with slice
	slice := []string{"satu", "dua", "tiga"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	//	for Range dengan slice
	for i, value := range slice {
		fmt.Println("Index", i, "with value:", value)
	}

	for _, value := range slice {
		fmt.Println("with value:", value)
	}

	// for with map
	person := make(map[string]string)
	person["name"] = "taufiqkba"
	person["title"] = "programmer"

	for key, value := range person {
		fmt.Println(key, ":", value)
	}
}
