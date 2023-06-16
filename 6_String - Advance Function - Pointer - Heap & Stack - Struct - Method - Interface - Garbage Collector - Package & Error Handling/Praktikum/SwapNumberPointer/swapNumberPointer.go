package main

import "fmt"

func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func main() {
	a := 10
	b := 20

	fmt.Println("Sebelum swap:")
	fmt.Println("a =", a)
	fmt.Println("b =", b)

	swap(&a, &b)

	fmt.Println("Setelah swap:")
	fmt.Println("a =", a)
	fmt.Println("b =", b)
}
