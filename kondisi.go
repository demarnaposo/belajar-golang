package main

import "fmt"

func main() {
	var nilai = 10

	switch nilai {
	case 10:
		fmt.Println("Sempurna")
	case 9:
		fmt.Println("Bagus")
	case 8:
		fmt.Println("Lumayan")
	default:
		fmt.Println("Belum Lulus")
	}

	// if nilai == 10 {
	// 	fmt.Println("lulus dengan sempurna")
	// } else if nilai > 7 {
	// 	fmt.Println("lulus")
	// } else if nilai == 6 {
	// 	fmt.Println("Sedikit Lagi Yukkk")
	// } else {
	// 	fmt.Println("Belajar Lagi")
	// }
}
