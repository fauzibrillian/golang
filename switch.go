package main

import "fmt"

func main() {
	nama := "fauzi"

	switch nama {
	case "eko":
		fmt.Println("halo eko")
		fmt.Println("halo eko")
	case "fauzi":
		fmt.Println("halo fauzi")
		fmt.Println("halo fauzi")
	default:
		fmt.Println("lah elu siapa ?")
		fmt.Println("lah elu siapa ?")
	}

	// nilai := 75
	// switch {
	// case nilai > 60:
	// 	fmt.Println("C")
	// case nilai > 70:
	// 	fmt.Println("B")
	// case nilai > 80:
	// 	fmt.Println("A")
	// }

	panjang := len(nama)
	switch {
	case panjang > 10:
		fmt.Println("terlalu panjang")
	case panjang > 5:
		fmt.Println("nama lumayan benar")
	default:
		fmt.Println("sudah benar")
	}

}
