package main

import "fmt"

func main() {
	slice := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	num := "0"
	fmt.Print("Digite um número: ")
	fmt.Scan(&num)
	for i, elemento := range slice {
		if elemento == num {
			slice = append(slice[:i], slice[i+1:]...)
		}
	}
	for _, elemento := range slice {
		println(elemento)
	}
}
