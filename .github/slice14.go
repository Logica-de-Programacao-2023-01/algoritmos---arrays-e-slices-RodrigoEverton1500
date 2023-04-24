package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	var x, y int
	fmt.Print("Digite o primeiro elemento: ")
	fmt.Scan(&x)
	fmt.Print("Digite o segundo elemento: ")
	fmt.Scan(&y)

	slice[x], slice[y] = slice[y], slice[x]

	fmt.Println(slice)
}
