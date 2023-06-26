package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	num := 0
	fmt.Print("Digite um número: ")
	fmt.Scan(&num)
	for _, elemento := range slice {
		if elemento == num {
			println("Número ", num, " está na lista")
		} else {
			slice = append(slice, num)
		}
	}
	slice = append(slice[:6], slice[10:]...)
	for _, elemento := range slice {
		println(elemento)
	}
}
