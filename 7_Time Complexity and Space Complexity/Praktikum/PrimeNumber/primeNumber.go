package main

import (
	"fmt"
	"math"
)

func PrimeNumber(number int) bool {
	if number <= 1 {
		return false
	}
	if number <= 3 {
		return true
	}

	if number%2 == 0 || number%3 == 0 {
		return false
	}

	sqrt := int(math.Sqrt(float64(number)))

	for i := 5; i <= sqrt; i += 6 {
		if number%i == 0 || number%(i+2) == 0 {
			return false
		}
	}

	return true
}

func main() {

	fmt.Println(PrimeNumber(1000000007))
	fmt.Println(PrimeNumber(35))

}
