package main

import "fmt"

func main() {
	var name1 = "Andi"
	var name2 = "Andi"
	var result bool = name1 == name2
	fmt.Println(result)

	val1 := 100
	val2 := 150
	var result2 bool = val1 > val2
	fmt.Println(result2)
}
