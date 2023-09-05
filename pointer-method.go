package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) married() {
	man.Name = "Mr. " + man.Name
	// fmt.Println("di method", man.Name)
}

func main() {
	fauzi := Man{"fauzi"}
	fauzi.married()

	fmt.Println(fauzi.Name)
}
