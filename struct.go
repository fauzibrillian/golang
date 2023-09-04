package main

import "fmt"

type Customer struct {
	Nama, Address string
	Age           int
}

func main() {
	var eko Customer

	eko.Nama = "Eko"
	eko.Address = "Jawa Timur"
	eko.Age = 23

	fmt.Println(eko.Nama)
	fmt.Println(eko.Address)
	fmt.Println(eko.Age)

	joko := Customer{
		Nama:    "Joko",
		Address: "Jawa Tengah",
		Age:     25,
	}
	fmt.Println(joko.Nama)
	fmt.Println(joko.Address)
	fmt.Println(joko.Age)

	budi := Customer{"Budi", "Jakarta", 25}
	fmt.Println(budi)

}
