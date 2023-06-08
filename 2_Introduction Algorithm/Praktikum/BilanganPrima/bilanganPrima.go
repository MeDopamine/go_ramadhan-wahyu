package main

import (
	"fmt"
	"math"
	"strconv"
)

func isPrimeNumber(number int) bool {
	if number <= 1 {
		return false
	}

	for i := 2; i <= int(math.Sqrt(float64(number))); i++ {
		if number%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	var input string
	fmt.Print("Masukkan bilangan: ")
	fmt.Scanln(&input)

	number, err := strconv.Atoi(input)
	if err == nil {
		if isPrimeNumber(number) {
			fmt.Println("Bilangan Prima")
		} else {
			fmt.Println("Bukan Bilangan Prima")
		}
	} else {
		fmt.Println("Input tidak valid")
	}
}
