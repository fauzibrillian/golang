package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"`
}

type Contoh struct {
	Name        string
	Description string
}

func IsValid(data interface{}) bool {
	t := reflect.TypeOf(data)
	for x := 0; x < t.NumField(); x++ {
		field := t.Field(x)
		if field.Tag.Get("required") == "true" {
			value := reflect.ValueOf(data).Field(x).Interface()
			if value == "" {
				return false
			}
		}
	}
	return true
}

func main() {
	sample := Sample{"Fauzi"}

	var sampleType reflect.Type = reflect.TypeOf(sample)

	fmt.Println(sampleType.NumField())
	fmt.Println(sampleType.Field(0).Name)
	fmt.Println(sampleType.Field(0).Tag.Get("required"))
	fmt.Println(sampleType.Field(0).Tag.Get("max"))

	sample.Name = ""
	fmt.Println(IsValid(sample))

	contoh := Contoh{"Fauzi", "Adi"}
	fmt.Println(contoh)
}
