package main

import "fmt"

func getgoodbye(nama string) string {
	return "Good Bye" + nama
}

func main() {
	goodbye := getgoodbye
	result := goodbye("fauzi")

	fmt.Println(goodbye("fauzi"))
	fmt.Println(result)

}
