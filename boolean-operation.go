package main

import "fmt"

func main() {
	name1 := "Andi"
	name2 := "Andi"

	if name1 == "Andi" || name2 == "Andi" {
		fmt.Println("Match")
	} else {
		fmt.Println("Not Match")
	}
}
