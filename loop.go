package main

import "fmt"

func main() {

	for counter := 1; counter < 10; counter++ {
		fmt.Println("perulangan ke ", counter)
	}

	slice := []string{
		"fauzi",
		"adi",
		"dimas",
		"alem",
		"izza",
		"mafud",
	}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for _, value := range slice {
		// fmt.Println("Index", i, "=", value)
		fmt.Println(value)
	}
	person := make(map[string]string)
	person["name"] = "fauzi"
	person["tittle"] = "programmer golang"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
