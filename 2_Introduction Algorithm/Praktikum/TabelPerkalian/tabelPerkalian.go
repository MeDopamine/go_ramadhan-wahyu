package main

import "fmt"

func printMultiplicationTable(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			fmt.Print(i * j)
			if j < n {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func main() {
	var input int
	fmt.Print("Masukkan nilai n: ")
	fmt.Scanln(&input)

	printMultiplicationTable(input)
}
