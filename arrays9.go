package main

import "fmt"

func main() {
	array := [6]int{1, 2, 3, 4, 5, 6}
	for i, elemento := range array {
		fmt.Print("Digite um número: ")
		fmt.Scan(&elemento)
		array[i] = elemento
	}
	for _, elemento := range array {
		print(elemento, " ")
	}
}
