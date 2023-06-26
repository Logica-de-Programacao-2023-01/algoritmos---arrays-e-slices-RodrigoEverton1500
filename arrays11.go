package main

import "fmt"

func main() {
	array := [2][3]int{{1, 2}, {1, 2, 3}}
	coluna1 := 0
	linha1 := 0
	fmt.Print("Digite linha: ")
	fmt.Scan(&linha1)
	fmt.Print("Digite coluna: ")
	fmt.Scan(&coluna1)
	for linha := range array {
		for coluna, elemento := range array[linha] {
			if coluna1 == coluna && linha1 == linha {
				println(elemento)
			}
		}
	}
}
