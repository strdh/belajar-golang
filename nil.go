package main

import "fmt"

func generateMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string {
			"name" : name,
		}
	}
}

func main() {
	nameMap := generateMap("Andi");
	fmt.Println(nameMap)
}