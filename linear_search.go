package main

import (
	"fmt"
	"os"
)

func main() {

	a := []int{2, 45, 896, 3}
	search := 3
	for pos, item := range a {
		if item == search {
			fmt.Println("Item found at position", pos)
			os.Exit(0)
		}
	}
	fmt.Print("Item not found")

}
