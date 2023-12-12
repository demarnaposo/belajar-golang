package main

import "fmt"

func main() {
	fmt.Println(tampilkan_pesan(123, 456))
}

func tampilkan_pesan(x int, y int) int {
	var z = x * y
	return z
}
