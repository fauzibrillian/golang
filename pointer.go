package main

import "fmt"

type Alamat struct {
	kota, provinsi, negara string
}

func ChangeCountryTo(address *Alamat) {
	address.negara = "Australia"
}

func main() {
	address1 := Alamat{"Sidoarjo", "Jawa Timur", "Indonesia"}
	address2 := &address1
	var address3 *Alamat = &address1

	address2.kota = "Surabaya"

	*address2 = Alamat{"Jakarta", "DKI", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	var address4 *Alamat = new(Alamat)
	address4.kota = "lampung"
	fmt.Println(address4)

	var Address = Alamat{
		kota:     "Malang",
		provinsi: "Jawa Timur",
		negara:   "",
	}

	ChangeCountryTo(&Address)
	fmt.Println(Address)
}
