package helper

import "fmt"

var version = 1
var Application = "hello golang"

func Hello(name string) {
	fmt.Println("Hellow jancog", name)
}

func goodbye(name string) {
	fmt.Println("Goodbye")
}
