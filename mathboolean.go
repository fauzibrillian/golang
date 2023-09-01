package main

import "fmt"

func main() {
	var ujian = 90
	const absensi = 80

	// var lulusujian = ujian >= 80
	// var lulusabsensi = absensi >= 70

	// fmt.Println(lulusabsensi)
	// fmt.Println(lulusujian)

	fmt.Println(ujian >= 80 && absensi >= 70)
}
