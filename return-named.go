package main

import "fmt"

func namalengkap() (firstname string, lastname string, age int) {
	firstname = "fauzi"
	lastname = "brillian"
	age = 23
	return
}

func main() {
	a, b, c := namalengkap()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
