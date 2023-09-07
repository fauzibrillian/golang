package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()

	data.PushBack("Fauzi")
	data.PushBack("Brillian")
	data.PushBack("Ananta")

	// fmt.Println(data.Front().Prev())
	// fmt.Println(data.Back().Next())

	// x adalaha data
	// belakang ke depan
	for x := data.Back(); x != nil; x = x.Prev() {
		fmt.Println(x.Value)
	}

	//depan ke belakang
	for x := data.Front(); x != nil; x = x.Next() {
		fmt.Println(x.Value)
	}
}
