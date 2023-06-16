package main

import "fmt"

func pow(x, n int) int {
	if n == 0 {
		return 1
	}

	halfPower := pow(x, n/2)

	if n%2 == 0 {
		return halfPower * halfPower
	}

	return x * halfPower * halfPower
}

func main() {
	fmt.Println(pow(2, 3))
	fmt.Println(pow(5, 3))
}
