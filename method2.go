package main

import (
	"fmt"
)

type gf struct {
	name string
	yb int
}

func (g gf) sayLove() {
	fmt.Println("I love you", g.name)
}

func (g * gf) setName(name string) {
	g.name = name
}


func main() {
	angie := gf{"Angie", 2022}
	angie.sayLove()
	angie.setName("Angie M")
	angie.sayLove()
}