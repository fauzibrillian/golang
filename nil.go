package main

import "fmt"

func newMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}

}

func main() {
	var orang map[string]string = newMap("fauzi")

	if orang == nil {
		fmt.Println("data kosong")
	} else {
		fmt.Println(orang)
	}

	// orang := newMap("Fauzi")
	// fmt.Println(orang)
}
