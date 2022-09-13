package main

import "fmt"

func main() {

	// if, else if, and else condition
	a := 21
	b := 20

	if a == b {
		fmt.Println(":", a+b)
	} else if a < b {
		fmt.Println(":", a-b)
	} else if a > b {
		fmt.Println(":", a*b)
	} else {
		fmt.Println(":", a/b)
	}

	name := "taufiqkba"
	//	 if with short statement
	if length := len(name); length > 5 {
		fmt.Println("the name is:", name)
	}
}
