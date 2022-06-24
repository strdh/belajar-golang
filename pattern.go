package main 

import "fmt"

func main() {
	n := 5

	// for i := 1; i<=n; i++ {
	// 	for j := i; j<=n; j++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println("")
	// }

	// fmt.Println("")

	// for i := 1; i<=n; i++ {
	// 	for j := 1; j<=i; j++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println("")
	// }

	// fmt.Println("")

	// for i := 1; i<=n; i++ {
	// 	for j := i; j<n; j++ {
	// 		fmt.Print(" ")
	// 	}

	// 	for j := 1; j<=i; j++ {
	// 		fmt.Print("*")
	// 	}

	// 	for j := 2; j<=i; j++ {
	// 		fmt.Print("*")
	// 	}

	// 	fmt.Println("")
	// }

	// fmt.Println("")

	// for i := 1; i<=n; i++ {
	// 	for j := 2; j<=i; j++ {
	// 		fmt.Print(" ")
	// 	}

	// 	for j := i; j<=n; j++ {
	// 		fmt.Print("*")
	// 	}

	// 	for j := i; j<n; j++ {
	// 		fmt.Print("*")
	// 	}

	// 	fmt.Println("")
	// }

	// fmt.Println("")

	for i := 1; i<=n; i++ {
		for j := 1; j<=n*2; j++ {
			if (i == 1 || i == n) || (j == 1 || j == n*2)  {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}

	fmt.Println("")

	for i := 1; i<=n; i++ {
		for j := 1; j<=n*2; j++ {
			if (i == 1 || i == n) || (j == 1 || j == n*2) {
				fmt.Print("* ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println("")
	}

	txt := "Programer"

	for i := 0; i<len(txt); i++ {
		for j, k := 0, len(txt) - 1; j<len(txt); j, k = j+1, k-1{
			if j == i || (i == 0 || i == len(txt) - 1) {
				fmt.Print(string(txt[j]))
			} else if k == i {
				fmt.Print(string(txt[k]))
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}

}