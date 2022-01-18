package main

import "fmt"

func main() {
	var names = [3]string{"Rudi", "Santi", "Budi"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
}
