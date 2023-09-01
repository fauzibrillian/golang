package main

import "fmt"

func main() {
	var bulan = [...]string{
		"januari",
		"febuari",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
	}

	var slice1 = bulan[5:11]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// slice1[2] = "halo deck"
	// fmt.Println(slice1)

	var slice2 = bulan[11:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "jono")
	fmt.Println(slice3)
	slice3[1] = "bukan januari"

	fmt.Println(slice3)
	fmt.Println(slice2)
	fmt.Println(bulan)

	newSlice := make([]string, 2, 5)

	newSlice[0] = "eko"
	newSlice[1] = "ngomong"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copySlice := make([]string, 1, cap(newSlice))

	copy(copySlice, newSlice)
	fmt.Println(copySlice)

}
