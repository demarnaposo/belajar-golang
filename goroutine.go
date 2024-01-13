package main

import (
	"fmt"
	"runtime"
)

func main() {

	runtime.GOMAXPROCS(2)

	go tampilkan_pesan(5, "Saya kirim")
	tampilkan_pesan(5, "saya kedua")

	var input string
	fmt.Scanln(&input)
}

func tampilkan_pesan(x int, pesan string) {
	for i := 0; i < x; i++ {
		fmt.Println((i + 1), pesan)
	}
}
