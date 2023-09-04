package main

import "fmt"

func gethello(nama string) string {
	if nama == "" {
		return "hello siapa ya ?"
	} else {
		return "hello " + nama
	}
}

func main() {
	result := gethello("fauzi")
	fmt.Println(result)

	fmt.Println(gethello(""))
}
