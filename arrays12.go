package main

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	slice := []int{}
	for _, elemento := range array {
		if elemento%3 == 0 {
			slice = append(slice, elemento)
		}
	}
	for _, elemento := range slice {
		print(elemento, " ")
	}
}
