package main

import (
	"fmt"
	// "reflect"
)

func main() {
	// var names []string;
	// names = append(names, "Rudi", "Santi", "Budi")

	// for i := 0; i < len(names); i++ {
	// 	fmt.Println(names[i])
	// }

	// fmt.Println("Type of names:", reflect.TypeOf(names))
	// fmt.Println("Type of names2:", reflect.TypeOf(names2))

	names := []string{"Rudi", "Santi", "Budi"}
	names = append(names, "Anselma")
	fmt.Println(names)
}