package main

import "fmt"

func main() {

	a := []int{1, 4, 8, 9, 12, 15}
	low := 0
	high := len(a) - 1
	search := 40
	found := false
	for low <= high {
		mid := (low + high) / 2
		if a[mid] == search {
			fmt.Println("Item found at ", mid)
			found = true
			break
		} else if search < a[mid] {
			high = mid - 1
		} else if search > a[mid] {
			low = mid + 1
		}

	}

	if !found {
		fmt.Println("Item not found")
	}

}
