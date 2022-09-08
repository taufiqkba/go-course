package main

import "fmt"

func main() {
	// type declaration
	type NoKTP string
	type Married bool

	var noKtpTaufiq NoKTP = "0912839123"
	var marriedStatus Married = false

	fmt.Println(noKtpTaufiq)
	fmt.Println(marriedStatus)
}
