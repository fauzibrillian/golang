package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println("Argument : ")
	fmt.Println(args)

	fmt.Println(args[1])
	fmt.Println(args[2])

	nama, err := os.Hostname()
	if err == nil {
		fmt.Println("Hostname : ", nama)
	} else {
		fmt.Println("Eror dengan alasan", err.Error())
	}

	// user := os.Getenv("Testing_User")
	// name := os.Getenv("Testing_Name")

	// fmt.Println(user)
	// fmt.Println(name)
}
