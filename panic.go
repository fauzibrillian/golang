package main

import "fmt"

func endapp() {
	pesan := recover() //recover
	if pesan != nil {
		fmt.Println("eror dengan pesan: ", pesan)
	}
	fmt.Println("Aplikasi Selesai")
}

func runapp(error bool) {
	defer endapp()
	if error {
		panic("Aplikasi Error")
	}
	fmt.Println("Aplikasi Jalan")
}

func main() {
	runapp(true)
	fmt.Println("eko")
}
