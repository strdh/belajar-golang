package main

import "fmt"

//defer func yang bisa dijadwalkan
func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApp1() {
	defer logging()
	fmt.Println("Run Application")
}

//panic function yang akan berjalan ketika terjadi error
func endApp() {
	fmt.Println("Aplikasi telah berakhir")
	message := recover()
	fmt.Println("Aplikasi Error", message)
}

func runApp2(error bool) {
	defer endApp()
	if error {
		panic("App error")
	}
	fmt.Println("Aplikasi berjalan")
}

func main() {
	// runApp1()
	runApp2(true)
}
