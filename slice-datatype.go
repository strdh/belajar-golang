package main

import "fmt"

func main() {
	var arr = [10]int{10, 100, 15, 7, 17, 18, 4, 5, 3, 8}
	var slice1 = arr[1:5]
	var slice2 = arr[0:]
	fmt.Println(slice1)
	fmt.Println(slice2)
	slice1[0] = 300
	fmt.Println(arr)
}
