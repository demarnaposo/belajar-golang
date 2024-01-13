package main

import "fmt"

func main() {
	var nomor int = 10
	nomor = 30

	var tas string = "hello madafaka"
	var alamat_tas *string = &tas

	var alamat_nomor *int = &nomor

	fmt.Println("nilai dari nomor : ", nomor)
	fmt.Println("kalimat dari : ", tas)
	fmt.Println("alamat variabel nomor : ", alamat_nomor)
	fmt.Println("alamat variabel tas : ", alamat_tas)
}
