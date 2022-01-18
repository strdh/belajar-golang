package main

import "fmt"
import "math"

type BangunDatar interface {
	luas() float64
	keliling() float64
}

type Persegi struct {
	sisi float64
}

type Lingkaran struct {
	diameter float64
}

func (persegi Persegi) luas() float64 {
	return math.Pow(persegi.sisi, 2)
}

func (persegi Persegi) keliling() float64 {
	return persegi.sisi * 4
}

func (lingkaran Lingkaran) getJari() float64 {
	return lingkaran.diameter / 2
}

func (lingkaran Lingkaran) luas() float64 {
	return math.Pi * math.Pow(lingkaran.getJari(), 2)
}

func (lingkaran Lingkaran) keliling() float64 {
	return math.Pi * lingkaran.diameter
}

//interface kosong untuk tipe data yang tidak pasti
func foo(i int) interface{} {
	if i == 0 {
		return "AHAHAHAHA"
	} else {
		return 10
	}
}

func main() {
	var bd1 BangunDatar
	bd1 = Persegi{10}

	var bd2 BangunDatar
	bd2 = Lingkaran{8}

	fmt.Println("Luas persegi     : ", bd1.luas())
	fmt.Println("Keliling persegi : ", bd1.keliling())
	fmt.Println("===========================")
	fmt.Println("Luas lingkaran     : ", bd2.luas())
	fmt.Println("Keliling lingkaran : ", bd2.keliling())

	fmt.Println("Foo : ", foo(1))
}