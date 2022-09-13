package main

import "fmt"

func main() {
	name := "dddddddd"

	switch name {
	case "Taufiq":
		fmt.Println("Hello Taufiq")
	case "taufiqkba":
		fmt.Println("HEllo taufiqkba")
	default:
		fmt.Println("Hello Not Found")
	}

	//	Switch case with Short statement
	switch length := len(name); length == 5 {
	case true:
		fmt.Println("name to loong")
	case false:
		fmt.Println("its right!")
	}

	//	Switch without condition
	length2 := len(name)
	switch {
	case length2 > 5:
		fmt.Println("Name to loong!!")
	case length2 == 5:
		fmt.Println("Name its Enough!")
	default:
		fmt.Println("Err!")
	}
}
