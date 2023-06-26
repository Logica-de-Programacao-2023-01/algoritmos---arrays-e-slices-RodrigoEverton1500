package main

import "sort"

func main() {
	array1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice1 := []int{}
	slice2 := []int{}
	d := 0
	for _, e := range array1 {
		slice1 = append(slice1, e)
		slice2 = append(slice2, e)
	}
	sort.Ints(slice1)
	for i, e := range slice1 {
		if slice2[i] != e {
			println("Array não está em ordem crescente")
			break
		} else {
			d++
		}
		if d == len(slice1) {
			println("Array em ordem crescente")
		}

	}
}
