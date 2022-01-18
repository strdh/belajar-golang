package main

import "fmt"

type Person struct {
	name, address string
	age           int
}

func (person Person) sayHi(name string) {
	fmt.Println("Hello", name, "My Name is", person.name)
}

func main() {
	var friend Person
	friend.name = "Mas Pono"
	friend.address = "Jakarta"
	friend.age = 20

	friend.sayHi("Budi")
	// fmt.Println(friend.name)
	// fmt.Println(friend.address)
	// fmt.Println(friend.age)

	// bt := Person{"Syahroni", "Bandung", 20}
	// fmt.Println(bt.name)
	// fmt.Println(bt.address)
	// fmt.Println(bt.age)

	// bt2 := Person{
	// 	name:    "Anita",
	// 	address: "Jakarta",
	// 	age:     22,
	// }
	// fmt.Println(bt2.name)
	// fmt.Println(bt2.address)
	// fmt.Println(bt2.age)

}
