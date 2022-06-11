package main

import (
	"fmt"
)

func sayHello(name string) string {
	return "Hello " + name
}

func add(a, b int) int {
	return a + b
}

//multiple return values
func foo1() (name1, name2 string) {
	name1 = sayHello("Kante")
	name2 = sayHello("Jorginho")
	return
}

//variadic function
func sumAll(numbers ...int) int {
	sum := 0
	for _, val := range numbers {
		sum += val
	}
	return sum
}

func main() {
	//closure
	getMinMax := func(numbers []int) (min, max int) {
		min = numbers[0]
		max = numbers[0]
		for _, val := range numbers {
			if val < min {
				min = val
			}
			if val > max {
				max = val
			}
		}
		return
	}

	// Immediately-Invoked Function Expression (IIFE)
	avg := func(numbers...int) int {
		sum := 0
		for _, val := range numbers {
			sum += val
		}
		return sum / len(numbers)
	}(10, 20, 30, 50, 100)

	//closure as return
	
	var min, max = getMinMax([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	// fmt.Println(sayHello("Kante"))
	fmt.Println(add(1, 2))
	fmt.Println(foo1())
	fmt.Println(sumAll(1, 2, 3, 4, 5))
	fmt.Println(getMinMax([]int{1, 2, 3, 4, 5}))
	fmt.Println(avg)
	fmt.Printf("Min : %v\nMax : %v", min, max)
}