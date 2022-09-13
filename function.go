package main

import "fmt"

func sayHello() {
	fmt.Println("say Hello!")
}

//function with parameter
func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello, ", firstName, lastName)
}

//function with return value
func add(a int, b int) int {
	if a%2 == 0 {
		return a * b
	} else {

	}
	return a + b
}

//function return with return value
func getHello(name string) string {
	return "Hi, " + name
}

//function return multiple values
func getFullName() (string, string, string) {
	return "Valorant", "Champions", "Istanbul"
}

//function with Named Return Values
func getCompleteAddress() (city string, country string, zipcode string) {
	city = "Semarang"
	country = "Indonesia"
	zipcode = "50249"
	return
}

func main() {
	//call a function
	sayHello()

	//call function with parameter
	sayHelloTo("taufiq", "kurniawan")

	//	function with return value
	resultAdd := add(6, 4)
	fmt.Println(resultAdd)

	returnResult := getHello("taufiqkbaaa")
	fmt.Println(returnResult)

	//	function return multiple values
	// using _ to ingnore a variable return values
	firstName, _, locationName := getFullName()
	fmt.Println(firstName)
	//fmt.Println(lastName)
	fmt.Println(locationName)

	//	function with Named Return Values
	city, _, zipcode := getCompleteAddress()
	fmt.Println(city)
	//fmt.Println(country)
	fmt.Println(zipcode)
}
