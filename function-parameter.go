package main

import "fmt"

type Filter func(string) string

func hellofilter(nama string, filter Filter) {
	namafiltered := filter(nama)
	fmt.Println("hello", namafiltered)
}

func spamfilter(nama string) string {
	if nama == "anjing" {
		return "..."
	} else {
		return nama
	}
}

func main() {
	hellofilter("fauzi", spamfilter)
	hellofilter("anjing", spamfilter)

	filter := spamfilter
	hellofilter("anjing", filter)
}
