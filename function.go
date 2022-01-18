package main

import (
	"fmt"
)

//multiple return
func foo1() (name1, name2 string) {
	name1 = "Kante"
	name2 = "Jorginho"
	return
}

//variadic function
func foo2(caption string, names ...string) {
	for i, val := range names {
		fmt.Println(caption, " ", i, " : ", val)
	}
}

//using arr / slice as parameter
func foo3(names []string) {
	for i, val := range names {
		fmt.Println("Anggota ", i+1, "", val)
	}
}

//func as val
func foo4(name string) string {
	return "Good Morning " + name
}

//func as paramter
func foo5(name string, filter func(string, []string) bool, fwords []string) {
	if filter(name, fwords) == true {
		fmt.Print("Hello ", name)
	} else {
		fmt.Print("Tidak lolos filter")
	}
}

func filter(name string, fwords []string) bool {
	for _, val := range fwords {
		if val == name {
			return false
		}
	}
	return true
}

//anonymous func
type Blacklist func(string) bool

func foo6(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked")
	} else {
		fmt.Println("Welcome home", name)
	}
}

//recursive func
func facto(n int) int {
	if n == 1 {
		return 1
	} else {
		return n * facto(n-1)
	}
}

func main() {
	// name1, name2 := foo1()
	// fmt.Print(name1, " ", name2)
	// foo2("Anggota", "Kante", "Hazard", "Mendy")
	// foo3([]string{"Kepa", "Pulisic", "Hazard"})
	// good := foo4("Andi")
	// fmt.Print(good)
	// fwords := []string{"Asu", "Celeng", "Gemblok"}
	// foo5("Asu", filter, fwords)
	// blacklist := func(name string) bool {
	// 	return name == "Andjing"
	// }
	// foo6("Andjing", blacklist)
	// foo6("Andi", blacklist)
	// fmt.Print(facto(5))
}
