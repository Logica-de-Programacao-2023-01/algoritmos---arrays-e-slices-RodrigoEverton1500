package main

func main() {
	array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice := []int{}
	for _, e := range array {
		if e%2 == 0 {
			slice = append(slice, e)
		}
	}
	for _, e := range slice {
		print(e, " ")
	}
}
