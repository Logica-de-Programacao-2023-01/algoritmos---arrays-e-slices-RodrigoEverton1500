package main

import "fmt"

func main() {
	array := [2][3]int{{}, {}}
	for linha := range array {
		for coluna, elemento := range array[linha] {
			println("Linha ", linha, " Coluna ", coluna)
			fmt.Println("Digite um elemento: ")
			fmt.Scan(&elemento)
			array[linha][coluna] = elemento
		}
	}
	for linha := range array {
		for coluna, elemento := range array[linha] {
			println("Linha ", linha, " Coluna ", coluna, " Elemento: ", elemento)
		}
	}
}
