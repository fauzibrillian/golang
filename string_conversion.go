package main

import (
	"fmt"
	"strconv"
)

func main() {
	boolean, err := strconv.ParseBool("true")

	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println(err.Error())
	}

	angka, err := strconv.ParseInt("10000", 10, 16)
	if err == nil {
		fmt.Println(angka)
	} else {
		fmt.Println(err.Error())
	}

	value := strconv.FormatInt(1000000, 16)
	fmt.Println(value)

	q := strconv.QuoteRuneToASCII('☺')
	fmt.Println(q)
}
