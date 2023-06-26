package main

import "sort"

func main() {
	slice := []int{1, 9, 8, 7, 6, 5, 4, 3, 2, 10}
	sort.Ints(slice)
	println("Maior ", slice[9], " Menor ", slice[0])
}
