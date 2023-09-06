package main

import (
	"flag"
	"fmt"
)

func main() {
	host := flag.String("host", "localhost", "Put your host")
	user := flag.String("user", "root", "Put your db user")
	password := flag.String("password", "root", "Put your db password")
	angka := flag.Int("angka", 0, "Put your angka")

	flag.Parse()

	fmt.Println("Host : ", *host)
	fmt.Println("User : ", *user)
	fmt.Println("Password : ", *password)
	fmt.Println("Angka : ", *angka)

}
