package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	total := sumAll(10, 20, 30, 40, 50, 60, 70, 80, 90)
	fmt.Println(total)

	// using slice to be parameter into variadic function
	slice := []int{33, 44, 55, 66}
	total = sumAll(slice...)
	fmt.Println(total)
}
