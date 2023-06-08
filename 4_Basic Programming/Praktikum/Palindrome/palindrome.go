package main

import "fmt"

func isPalindrome(str string) bool {
	length := len(str)
	for i := 0; i < length/2; i++ {
		if str[i] != str[length-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var input string
	fmt.Print("Masukkan string: ")
	fmt.Scanln(&input)

	if isPalindrome(input) {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Bukan palindrome")
	}
}
