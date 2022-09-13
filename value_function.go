package main

import "fmt"

func getGoodBye(name string) string {
	return "Bye... " + name
}

func main() {
	getBye := getGoodBye
	fmt.Println(getBye("taufiqkba"))

}
