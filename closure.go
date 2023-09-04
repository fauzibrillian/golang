package main

import "fmt"

func main() {
	name := "eko"
	counter := 0

	increment := func() {
		name = "budi"
		fmt.Println("increment")
		counter++
	}

	increment()
	increment()
	fmt.Println(counter)
	fmt.Println(name)
}
