package main

import "fmt"

func PairSum(array []int, target int) []int {
	left := 0
	right := len(array) - 1

	for left < right {
		sum := array[left] + array[right]

		if sum == target {
			return []int{left, right}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}

	return []int{} // Return an empty slice if no pair is found
}

func main() {
	fmt.Println(PairSum([]int{1, 2, 3, 4, 6}, 5)) // Output: [1, 3]
	fmt.Println(PairSum([]int{2, 5, 9, 11}, 11))  // Output: [0, 2]
}
