package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Fauzi Brillian", "Fauzi"))
	fmt.Println(strings.Contains("Fauzi Brillian", "Jono"))

	fmt.Println(strings.Split("Fauzi Brillian", " "))

	fmt.Println(strings.ToLower("Fauzi Brillian Ananta"))
	fmt.Println(strings.ToUpper("Fauzi Brillian Ananta"))

	fmt.Println(strings.Trim("      Fauzi Brillian         ", " "))

	fmt.Println(strings.ReplaceAll("Fauzi Fauzi Fauzi", "Jono", "Joko"))
}
