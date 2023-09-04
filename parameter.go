package main

import "fmt"

func helloto(firstname string, lastname string, angka int) {
	fmt.Println("hello", firstname, lastname, angka)
}

func main() {
	firstname := "eko"
	angka := 20
	helloto("fauzi", "brillian", angka)
	helloto(firstname, "kurniawan", angka)
}
