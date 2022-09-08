package main

import "fmt"

func main() {
	// conversion data type
	var nilai32 int32 = 129
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8) // back to minimum value (be careful if want to change data type)

	// conversion data type to string
	var name = "Taufiq Kurniawan"
	var t = name[0]
	var tString string = string(t)

	fmt.Println(name)
	fmt.Println(tString)
}
