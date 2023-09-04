package main

import "fmt"

type Blacklist func(string) bool

func register(nama string, blacklist Blacklist) {
	if blacklist(nama) {
		fmt.Println("you are blocked", nama)
	} else {
		fmt.Println("Welcome", nama)
	}
}

func main() {
	blacklist := func(nama string) bool {
		return nama == "admin"
	}

	register("admin", blacklist)
	register("eko", blacklist)

	register("root", func(nama string) bool {
		return nama == "root"
	})

	register("eko", func(nama string) bool {
		return nama == "root"
	})
}
