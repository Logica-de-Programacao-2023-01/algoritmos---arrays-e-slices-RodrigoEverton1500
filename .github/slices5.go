package main

import "fmt"

func main() {
	matriz := [3][2]int{}
	fmt.Println(matriz)
	for i := 0; i < matriz[2][1]; i++ {
		break
	}
}
