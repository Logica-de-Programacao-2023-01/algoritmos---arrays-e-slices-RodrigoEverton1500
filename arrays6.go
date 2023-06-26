package main

import "fmt"

func main() {
	array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	num := int(0)
	for {
		fmt.Print("Digite um número: ")
		fmt.Scan(&num)
		for _, elemento := range array {
			if num == elemento {
				println("Número ", elemento, " encontrado")
				continue
			} else {
				continue
			}
		}
	}

}
