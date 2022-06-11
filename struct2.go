package main

import (
	"fmt"
)

type student struct {
	name string
	yb int
}


func main() {
	// var students []student
	// students = append(students, student{"Rudi", 1998}, student{"Santi", 1999}, student{"Budi", 2000})

	var students = []student{
		{"Rudi", 1998},
		{"Santi", 1999},
		{"Budi", 2000},
	}

	for i := 0; i < len(students); i++ {
		fmt.Println("Student ", (i+1), "\nName: ", students[i].name, "\nBirth Year", students[i].yb, "\n")
	}

	var gf = struct {
		name string
		city string
	}{
		name: "Anselma",
		city: "Bandung",
	}

	fmt.Println("\n\nGirlfriend: ", gf.name, "\nCity: ", gf.city)

	// s1 := student{"Rudi", 1998}
	// s2 := s1 
	// fmt.Println(s2.name)
	
	// s1 := student{"Anselma", 2002}
	// sc1 := scout{"A", s1}
	// fmt.Println(sc1.student.name)


}