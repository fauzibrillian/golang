package main

import (
	"fmt"
)

func main() {
	var a = 10
	var b = 20
	var c = a * b

	var x = 12
	x += 12
	x++

	const nama1 = "fauzi"
	const nama2 = "jono"

	var result bool = nama1 != nama2

	fmt.Println(result)
	fmt.Println(c)
	fmt.Println(x)
}
