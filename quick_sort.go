package main

import "fmt"

var a = []int{0, 5, 2, 1, 6, 3, -1, 4, 5, 8, 9, 2}

func main() {

	leftPointer := 0
	rightPointer := len(a) - 1
	fmt.Println("Unsorted Array is ", a)
	quickSort(leftPointer, rightPointer)
	fmt.Println("Sorted Array is", a)
}

func quickSort(leftPointer, rightPointer int) {

	/* The Quicksort algorithm relies heavily on partitions. It works as follows:

	   Partition the array. The pivot is now in its proper place.

	   Treat the subarrays to the left and right of the pivot as their own arrays, and recursively repeat steps #1 and #2. That means that we’ll partition each subarray, and end up with even smaller subarrays to the left and right of each subarray’s pivot. We then partition those subarrays, and so on and so forth.

	   When we have a subarray that has zero or one elements, that is our base case and we do nothing.
	*/
	
	if rightPointer-leftPointer <= 0 {
		return
	}

	lastPivotPoistion := partition(leftPointer, rightPointer)
	quickSort(leftPointer, lastPivotPoistion-1)
	quickSort(lastPivotPoistion+1, rightPointer)

}

func partition(leftPointer, rightPointer int) int {
	pivotPoistion := rightPointer
	pivotValue := a[pivotPoistion]
	rightPointer = rightPointer - 1
	fmt.Println("Right pointer", rightPointer)
	for {

		for {
			if a[leftPointer] < pivotValue {
				leftPointer++
			} else {
				fmt.Println("Break Left Pointer", leftPointer)
				break
			}
		}
		for {

			if a[rightPointer] > pivotValue {
				if rightPointer >= 1 {
					rightPointer--
				} else {

					fmt.Println("Right pointer", rightPointer)
					break
				}

			} else {
				fmt.Println("Break Right Pointer", rightPointer)
				break
			}
		}
		if leftPointer >= rightPointer {
			break
		} else {
			swap(leftPointer, rightPointer)
		}

	}
	swap(leftPointer, pivotPoistion)
	return leftPointer

}

func swap(lp, rp int) {
	t := a[lp]
	a[lp] = a[rp]
	a[rp] = t

}
