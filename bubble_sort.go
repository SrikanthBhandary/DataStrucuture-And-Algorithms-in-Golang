package main

import "fmt"

func main() {
	a := []int{2, 3, 4, 5, 1, -1}
	n := len(a) - 1
	fmt.Println("Unsorted Array is ", a)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if a[j] > a[j+1] {
				t := a[j+1]
				a[j+1] = a[j]
				a[j] = t
			}
		}

	}
	fmt.Println("Sorted Array is ", a)

}
