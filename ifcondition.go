package main

import "fmt"

func main() {
	nama := "muhammad indra syahfiti"

	if nama == "fauzi" {
		fmt.Println("halo fauzi")
	} else if nama == "joko" {
		fmt.Println("halo joko")
	} else {
		fmt.Println("lah siapa ?")
	}

	if panjang := len(nama); panjang > 5 {
		fmt.Println("terlalu panjang")
	} else {
		fmt.Println(nama)
	}

}
