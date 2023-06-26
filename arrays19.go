package main

func main() {
	array1 := [5]int{1, 2, 3, 4, 5}
	array2 := [5]int{10, 9, 8, 7, 6}
	n := 0
	for _, e := range array1 {
		n += e
	}
	for _, e := range array2 {
		n += e
	}
	println(n)
}
