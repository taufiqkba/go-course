package main

import "fmt"

func main() {

	finalScore := 79
	present := 80

	var lulusNilaiAkhir = finalScore > 80
	var lulusPresent = present > 75

	var lulus = lulusNilaiAkhir && lulusPresent
	fmt.Println(lulus)
	fmt.Println(finalScore >= 80 && present >= 80)
}
