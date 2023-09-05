package main

import "fmt"

func random() interface{} {
	return "eko"
}

func main() {
	var result interface{} = random()
	// var resultString string = result.(string)
	// fmt.Println(resultString)

	switch hasil := result.(type) {
	case string:
		fmt.Println("value", hasil, "adalah string")
	case int:
		fmt.Println("value", hasil, "adalah int")
	default:
		fmt.Println("unknown type")
	}
}
