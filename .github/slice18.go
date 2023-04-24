package main

import "fmt"

func main() {
	var n int
	fmt.Print("Digite um número: ")
	fmt.Scan(&n)

	var primos []int

	for i := 2; i <= n; i++ {
		primo := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				primo = false
				break
			}
		}
		if primo {
			primos = append(primos, i)
		}
	}

	fmt.Println(primos)
}
