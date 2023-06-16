package main

import "fmt"

func arrayContains(arr []string, name string) bool {
	for _, value := range arr {
		if value == name {
			return true
		}
	}
	return false
}

func ArrayMerge(arrayA, arrayB []string) []string {
	merged := make([]string, len(arrayA))
	copy(merged, arrayA)

	for _, name := range arrayB {
		if !arrayContains(merged, name) {
			merged = append(merged, name)
		}
	}

	return merged
}

func main() {
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))

	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))

}
