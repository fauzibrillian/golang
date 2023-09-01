package main

import "fmt"

func main() {
	type noktp string
	type nikah bool

	var nofauzi noktp = "624634666"
	const statusnikah nikah = false
	fmt.Println(nofauzi)
	fmt.Println(statusnikah)
}
