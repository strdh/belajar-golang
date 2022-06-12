package main

import (
	f "fmt"
)

type person struct {
	name string
	yb int
}

func main() {
	// var any interface{} = person{name:"Ansel", yb:2022}
	// name := any.(person).name
	// f.Println(name)

	// any = "Angie Marcheria"
	// f.Println(any)

	// any = []string{"Anselma Hanadya Putri", "Angie Marcheria", "Yolanda Tamara"}
	// f.Println(any)

	// any = 12.5
	// f.Println(any)

	// result := any.(float64) * 10
	// f.Println(result)

	// var data map[string]interface{}
	// data = map[string]interface{}{
	// 	"name" : "Angie",
	// 	"yb" : 2022,
	// }

	// f.Println(data["name"])

	// var data map[string]any

	// data = map[string]any{
	// 	"name" : "Angie",
	// 	"hobbies" : []string{"Eat", "Selfie"},
	// }

	// f.Println(data["hobbies"])

	var person = []map[string]interface{}{
		{"name" : "Anselma", "age" : 20},
		{"name" : "Angie", "age" : 20},
	}

	for _, p := range person{
		f.Printf("Name : %v | Age : %v\n", p["name"], p["age"])
	}
}