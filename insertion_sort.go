/*
Selection Sort or Insertion Sort?
The answer is: well, it depends. In an average case—where an array is randomly sorted—they perform similarly.
If you have reason to assume that you’ll be dealing with data that is mostly sorted, Insertion Sort will be a better choice.
If you have reason to assume that you’ll be dealing with data that is mostly sorted in reverse order, Selection Sort will be faster.

	If you have no idea what the data will be like, that’s essentially an average case, and both will be equal.
*/
package main

import "fmt"

func main() {

	a := []int{4, 5, 1, 0, 7, -1, 0}
	n := len(a)
	fmt.Println("Unsorted array", a)
	runs := 0
	for i := 1; i < n; i++ {
		pos := i
		tempValue := a[pos]

		for pos > 0 && a[pos-1] > tempValue {
			a[pos] = a[pos-1]
			pos--
			runs++
		}
		a[pos] = tempValue

	}
	fmt.Println("Sorted Array", a)
	fmt.Println("Total runs", runs)
}
