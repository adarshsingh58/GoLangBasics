package main

import (
	"math/rand"
	"time"
	"fmt"
)

func main() {
	arr := generateRandomSliceQuickSort()
	fmt.Println(arr)
	arr = quickSort(arr)
	fmt.Println(arr)
}
func quickSort(arr []int) []int {

	partitionIndex := findPartition(arr)
	if partitionIndex >0 {
		quickSort(arr[:partitionIndex])
		if partitionIndex == len(arr) {
			quickSort(arr[partitionIndex:])
		} else {
			quickSort(arr[partitionIndex+1:])
		}
	}
	return arr
}
func findPartition(arr []int) int {
	pivotIndex := len(arr) - 1
	partitionIndex := 0
	if len(arr) > 2 {
		fmt.Println(arr, " pivot ", arr[pivotIndex], " paertitionInd", partitionIndex)
		for i, v := range arr {
			if partitionIndex < len(arr) && i < len(arr) && v > arr[pivotIndex] {
				temp := arr[partitionIndex]
				arr[partitionIndex] = arr[i]
				arr[i] = temp
				partitionIndex++
			}
		}
	}
	temp := arr[partitionIndex]
	arr[partitionIndex] = arr[pivotIndex]
	arr[pivotIndex] = temp
	fmt.Println(arr, " pivot ", arr[pivotIndex], " paertitionInd", partitionIndex)
	return partitionIndex
}

func generateRandomSliceQuickSort() ([]int) {
	slice := make([]int, 10, 20)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		slice[i] = rand.Intn(100)
	}
	return slice
}
