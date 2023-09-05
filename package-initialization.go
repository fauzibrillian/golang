package main

import (
	"belajar_go/database"
	"fmt"
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}
