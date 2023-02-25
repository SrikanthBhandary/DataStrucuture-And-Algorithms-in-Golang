/*
We check each cell of the array from left to right to determine which value is least.
As we move from cell to cell, we keep in a variable the lowest value we’ve encountered so far.

If we encounter a cell that contains a value that is even less than the one in our variable, we replace it so that the variable now points to the new index.

Once we’ve determined which index contains the lowest value, we swap that index with the value we began the passthrough with.
This would be index 0 in the first passthrough, index 1 in the second passthrough, and so on and so forth.
Repeat steps 1 and 2 until all the data is sorted.

*/

package main

import "fmt"

func main() {
	a := []int{1, -1, 0, 4, 2, 5, 7, 8, 9, 3, 8, 2, 6, 8, 2, 8, 2, 8, 52, 8, 5, 8, 52, 85, -100}
	n := len(a) - 1

	for i := 0; i < n; i++ {
		minPos := i
		minValue := a[i]
		for j := i + 1; j <= n; j++ {
			if a[j] < minValue {
				minPos = j
				minValue = a[j]
			}

		}
		t := a[i]
		a[i] = a[minPos]
		a[minPos] = t

	}
	fmt.Println(a)

}
