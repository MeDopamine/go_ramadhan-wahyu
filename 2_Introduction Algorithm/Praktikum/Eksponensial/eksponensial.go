package main

import "fmt"

func power(x float64, n int) float64 {
	result := 1.0

	if n >= 0 {
		for i := 0; i < n; i++ {
			result *= x
		}
	} else {
		for i := 0; i > n; i-- {
			result /= x
		}
	}

	return result
}

func main() {
	var x float64
	var n int

	fmt.Print("Enter the value of x: ")
	fmt.Scanln(&x)

	fmt.Print("Enter the value of n: ")
	fmt.Scanln(&n)

	result := power(x, n)
	fmt.Println("Result:", result)
}
