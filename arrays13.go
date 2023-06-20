package main

import "fmt"

func main() {
	array := [7]int{1, 2, 3, 4, 5, 6, 7}
	num1 := 0
	num2 := 0
	fmt.Print("Digite primeiro número: ")
	fmt.Scan(&num1)
	fmt.Print("Digite segundo número: ")
	fmt.Scan(&num2)
	array[0] = num1
	array[6] = num2
	for _, e := range array {
		print(e, " ")
	}
}
