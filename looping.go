package main

import "fmt"

func main() {
	// n := 5
	// for i := 1; i <= n; i++ {
	// 	for j := 1; j <= i; j++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println("")
	// }
	names := []string{"Squidword", "Spongebob", "Mr Crab"}
	for i, val := range names {
		fmt.Println("Nama ke ", i+1, " : ", val)
	}
}
