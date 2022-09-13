package main

import "fmt"

func main() {
	//break
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println("looping into-", i)
	}
	//	continue
	fmt.Println("")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("looping into-", i)
	}
}
