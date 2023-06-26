package main

func main() {
	array := [5]float32{3.0, 4.0, 5.0, 6.0, 7.0}
	slice := []float32{}
	num := float32(0)
	for i, e := range array {
		num = array[i]
		if num > 5 {
			slice = append(slice, e)
		}
	}
	for _, e := range slice {
		print(e, " ")
	}
}
