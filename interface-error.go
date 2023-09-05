package main

import (
	"errors"
	"fmt"
)

func pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Engga bisa bro")
	} else {
		result := nilai / pembagi
		return result, nil
	}
}

func main() {
	var contoheror error = errors.New("Ups Eror bro")
	fmt.Println(contoheror.Error())

	hasil, err := pembagian(100, 0)
	if err == nil {
		fmt.Println("Hasil", hasil)
	} else {
		fmt.Println("Error", err.Error())
	}
}
