package main

import "fmt"

func main() {
	slice := []int{}
	var soma, tamanho, x int

	fmt.Println("Digite tamanho da slice: ")
	fmt.Scan(&tamanho)

	for i := 0; i < tamanho; i++ {
		fmt.Println("Digite um valor: ")
		fmt.Scan(&x)
		slice = append(slice, x)
		soma += x
	}
	fmt.Println(slice)
	fmt.Println(soma)
}
