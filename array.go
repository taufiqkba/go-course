package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "taufiq"
	names[1] = "kurniawan"
	names[2] = "bayu"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// another example array modification

	var values = [3]int{
		90, 95, 80,
	}

	fmt.Println(values)
	fmt.Println(len(values))

	values[0] = 1000
	fmt.Println(values)
	fmt.Println(values[2])

	var values2 [10]int

	fmt.Println(len(values2))
}
