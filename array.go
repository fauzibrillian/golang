package main

import "fmt"

func main() {

	var names [3]string
	names[0] = "fauzi"
	names[1] = "adi"
	names[2] = "tejo"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [4]int{
		100,
		20,
		30,
		100,
	}

	fmt.Println(values)

	var string = [3]string{
		"adi",
		"tejo",
		"jono",
	}

	fmt.Println(string)

}
