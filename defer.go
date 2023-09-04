package main

import "fmt"

func logging() {
	fmt.Println("selesai memanggil function")
}

func runapplication(value int) {
	defer logging()
	fmt.Println("run application")
	result := 10 / value
	fmt.Println("Result", result)
}

func main() {
	runapplication(10)
}
