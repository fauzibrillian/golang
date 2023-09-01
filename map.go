package main

import "fmt"

func main() {

	orang := map[string]string{
		"nama":   "fauzi",
		"tempat": "surabaya",
	}

	orang["jabatan"] = "programmer"

	fmt.Println(orang)
	fmt.Println(orang["nama"])
	fmt.Println(orang["tempat"])
	fmt.Println(orang["jabatan"])

	var buku map[string]string = make(map[string]string)
	buku["judul"] = "belajar golang"
	buku["author"] = "Eko"
	buku["ops"] = "salah"

	fmt.Println(buku)
	fmt.Println(len(buku))

	delete(buku, "ops")

	fmt.Println(buku)
	fmt.Println(len(buku))

}
