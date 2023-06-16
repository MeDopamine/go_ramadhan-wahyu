package main

import "fmt"

func getMinMax(numbers ...*int) (min int, max int) {
	min = *numbers[0]
	max = *numbers[0]

	for _, num := range numbers {
		if *num < min {
			min = *num
		}
		if *num > max {
			max = *num
		}
	}

	return min, max
}

func main() {
	var a1, a2, a3, a4, a5, a6 int
	fmt.Print("masukan angka: ")
	fmt.Scan(&a1, &a2, &a3, &a4, &a5, &a6)
	min, max := getMinMax(&a1, &a2, &a3, &a4, &a5, &a6)
	fmt.Println("nilai min: ", min)
	fmt.Println("nilai max: ", max)
}
