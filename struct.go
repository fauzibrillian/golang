package main

import "fmt"

type Customer struct {
	Nama, Address string
	Age           int
}

func (customer Customer) sayHi(Nama string) {
	fmt.Println("hello", Nama, "my name is ", customer.Nama)
}

func (a Customer) sayHu() {
	fmt.Println("Huuuu dari", a.Nama)
}

func main() {
	var eko Customer

	eko.Nama = "Eko"
	eko.Address = "Jawa Timur"
	eko.Age = 23

	eko.sayHi("Joko")
	eko.sayHu()

	// fmt.Println(eko.Nama)
	// fmt.Println(eko.Address)
	// fmt.Println(eko.Age)

	// joko := Customer{
	// 	Nama:    "Joko",
	// 	Address: "Jawa Tengah",
	// 	Age:     25,
	// }
	// fmt.Println(joko.Nama)
	// fmt.Println(joko.Address)
	// fmt.Println(joko.Age)

	// budi := Customer{"Budi", "Jakarta", 25}
	// fmt.Println(budi)

}
