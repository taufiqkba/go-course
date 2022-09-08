package main

import "fmt"

func main() {
	months := [...]string{
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	}

	slice1 := months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	months[5] = "Changed"
	fmt.Println(slice1)

	slice1[0] = "MayLagi"
	fmt.Println(slice1)

	//	Append Slice
	slice2 := months[2:4]
	fmt.Println(slice2)

	slice3 := append(slice2, "Taufiqkba")
	fmt.Println(slice3)

	slice3[1] = "Not December"
	fmt.Println(slice3)
	fmt.Println(slice2)
	fmt.Println(months)

	//	Make slice
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Taufiq"
	newSlice[1] = "Kurniawan"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	//	Copy Slice
	copySlice := make([]string, 1, cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	//	Array vs Slice
	thisIsArray := [5]int{1, 2, 3, 4, 5}
	thisIsSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(thisIsArray)
	fmt.Println(thisIsSlice)
}
