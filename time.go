package main

import (
	"fmt"
	"time"
)

func main() {
	sekarang := time.Now()
	fmt.Println(sekarang)
	fmt.Println(sekarang.Year())
	fmt.Println(sekarang.Month())
	fmt.Println(sekarang.Hour())
	fmt.Println(sekarang.Minute())
	fmt.Println(sekarang.Second())

	utc := time.Date(2023, 10, 10, 10, 10, 10, 10, time.UTC)
	fmt.Println(utc)

	layout := "2023-07-23"
	parse, _ := time.Parse(layout, "2020-07-06")
	fmt.Println(parse)
}
