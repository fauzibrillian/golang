package main

import "fmt"

func factorialloop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i //result = result * i
	}
	return result
}

func factorialrecursive(angka int) int {
	if angka == 1 {
		return 1
	} else {
		return angka * factorialrecursive(angka-1)
	}
}

func main() {
	loop := factorialloop(10)
	recursive := factorialrecursive(10)
	fmt.Println(loop)
	fmt.Println(recursive)
	// fmt.Println(5 * 4 * 3 * 2 * 1)
}
