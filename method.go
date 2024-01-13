package main

import "fmt"

func main() {

	var s1 pelajar
	var s2 = pelajar{"dems", 26}
	s1.nama = "naps"
	s1.umur = 20
	s1.namasaya()
	s2.namasaya()

}

type pelajar struct {
	nama string
	umur int
}

func (s pelajar) namasaya() {
	fmt.Println("nama saya adalah", s.nama)
	fmt.Println("umur saya : ", s.umur)

}
