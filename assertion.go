package main

import "fmt"

func foo(tipe string) interface{} {
	if tipe == "s" {
		return "STR"
	} else if tipe == "i" {
		return 100
	} else {
		return true
	}
}

func main() {
	var result interface{} = foo("s")
	// fmt.Println(result.(string))
	switch val := result.(type) {
	case string:
		fmt.Println("String", val)
	case int:
		fmt.Println("Int", val)
	default:
		fmt.Println("Uknown")
	}
}