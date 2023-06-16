package main

import (
	"fmt"
)

func Compare(A, B string) string {
	m := make(map[rune]bool)

	for _, char := range A {
		m[char] = true
	}

	common := ""
	for _, char := range B {
		if m[char] {
			common += string(char)
		}
	}

	return common
}

func main() {

	fmt.Println(Compare("aka", "akashi"))

	fmt.Println(Compare("kangooro", "kang"))
}
