package main

import "fmt"

func main() {
	var slice []int
	var x, y, z int
	fmt.Println("Digite o tamanho da slice: ")
	fmt.Scan(&y)
	for i := 0; i < y; i++ {
		fmt.Print("Digite um número: ")
		fmt.Scan(&x)
		slice = append(slice, x)
		x = x + z
	}
	fmt.Println(slice)
	fmt.Println(x)

}
