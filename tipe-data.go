package main

import "fmt"

func main() {
	// var number_positif uint8 = 10 // 0-255
	var number_positif uint16 = 10 // 0-65535
	// var number_positif uint32 = 10 //0-4294967295
	// var number_positif uint64 = 10 //0-18446744073709551615

	var number_negatif int64 = -99

	var apakah_ada bool = true
	var pesan string = "Halo teman"

	fmt.Println(number_positif)
	fmt.Println(number_negatif)
	fmt.Println(apakah_ada)
	fmt.Println(pesan)

}
