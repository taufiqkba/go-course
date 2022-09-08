package main

import "fmt"

func main() {

	person := map[string]string{
		"name":    "taufiq",
		"address": "pucang gading",
	}

	person["title"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	var book map[string]string = make(map[string]string)

	book["title"] = "Learn GoLang"
	book["author"] = "taufiq kurniawan"
	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "author")
	fmt.Println(book)
	fmt.Println(len(book))
}
