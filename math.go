package main

import "fmt"

func main() {

	// operate add
	a := 10
	b := 20
	result := a + b
	fmt.Println(result)

	// augmented assignment
	var c = 11
	c += 10
	fmt.Println(c)

	c -= 2
	fmt.Println(c)

	c *= 2
	fmt.Println(c)

	c /= 2
	fmt.Println(c)

	// unary operator
	d := 1
	d++
	fmt.Println(d)

	negative := -100
	fmt.Println(negative)
}
