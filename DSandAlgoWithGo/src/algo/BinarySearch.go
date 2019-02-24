package main

import "fmt"

func main() {
	arr := []int{12, 15, 16, 21, 26, 29, 52}
	op:=binarySearch(arr,21)
	fmt.Println(op)
}

func binarySearch(arr []int, target int) int {

	return binarySearchRecursive(arr, 0, len(arr)-1, target)
}
func binarySearchRecursive(arr []int, start int, end int, target int) (index int) {
	mid := (start + end) / 2
	if arr[mid] == target {
		return mid
	}
	if start <= end {
		if arr[mid] > target {
			return binarySearchRecursive(arr, start, mid, target)
		} else {
			return binarySearchRecursive(arr, mid, end, target)
		}

	} else {
		return -1
	}
}
