package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0

	for _, value := range numbers {
		total += value
	}
	return total
}

func main() {
	total := sumAll(80, 90, 70, 60, 50, 40)
	fmt.Println(total)

	slice := []int{80, 90, 70, 60, 50, 40}
	total = sumAll(slice...)
	fmt.Println(total)
}
