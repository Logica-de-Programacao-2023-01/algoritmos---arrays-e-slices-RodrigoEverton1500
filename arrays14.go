package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	num1 := 0
	num2 := 0
	fmt.Print("Digite o primeiro índice: ")
	fmt.Scan(&num1)
	fmt.Print("Digite o segundo índice: ")
	fmt.Scan(&num2)
	n1 := slice[num1]
	n2 := slice[num2]
	slice[num1] = n2
	slice[num2] = n1
	for _, elementos := range slice {
		fmt.Print(elementos, " ")
	}
}
