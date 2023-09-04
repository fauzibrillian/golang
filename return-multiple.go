package main

import "fmt"

func getnama() (string, string, int) {
	return "fauzi", "brillian", 23
}

func main() {
	firstname, lastname, _ := getnama()
	fmt.Println(firstname)
	fmt.Println(lastname)
	// fmt.Println(age)
}
