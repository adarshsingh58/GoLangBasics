package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	arr := GenerateRandomSlice_BubbleSort()
	fmt.Println(arr)
	arr = bubbleSort(arr)
	fmt.Println(arr)
}

func bubbleSort(arr []int) []int {
	for i := len(arr) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
	return arr
}

func GenerateRandomSlice_BubbleSort() ([]int) {
	slice := make([]int, 10, 20)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		slice[i]=rand.Intn(100)
	}
	return slice
}