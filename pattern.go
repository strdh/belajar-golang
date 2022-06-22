package main 

import "fmt"

func main() {
	n := 5
	for i := 1; i<=n; i++ {
		for j := i; j<=n; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}