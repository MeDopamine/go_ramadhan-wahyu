package main

import (
	"fmt"
	"strconv"
)

func munculSekali(angka string) []int {
	counter := make(map[int]int)

	for _, char := range angka {
		digit, _ := strconv.Atoi(string(char))
		counter[digit]++
	}

	result := []int{}
	for digit, count := range counter {
		if count == 1 {
			result = append(result, digit)
		}
	}

	return result
}

func main() {
	fmt.Println(munculSekali("1234123"))    // Output: [4]
	fmt.Println(munculSekali("12345"))      // Output: [1 2 3 4 5]
	fmt.Println(munculSekali("1122334455")) // Output: []
	fmt.Println(munculSekali("76523752"))   // Output: [6 3]
}
