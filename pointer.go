package main

import "fmt"

func swap(a *int, b *int) {
	temp := *a 
	*a = *b 
	*b = temp
}

func main() {
	// angka1 := 100
	// addr := &angka1
	// fmt.Println(*addr)
	// *addr = 1000
	// fmt.Println(*addr)
	a := 100
	b := 10
	swap(&a, &b)
	fmt.Println("nilai a : ", a)
	fmt.Println("nilai b : ", b)
}