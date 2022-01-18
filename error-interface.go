package main

import (
	"fmt"
	"errors"
)

func bagi(val float64, div float64) (float64, error) {
	if div == 0 {
		return 0, errors.New("Pembagi tidak boleh 0")
	} else {
		return val / div,nil
	}
}

func main() {
	fmt.Println(bagi(10, 0))
	fmt.Println(bagi(100, 5))
}