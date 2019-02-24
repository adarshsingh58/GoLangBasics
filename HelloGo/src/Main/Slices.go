package main

import "fmt"

func main() {
	creatingSlice()
	creatingSliceShortHand()
	sliceASlice()
	increasingSizeOfSlice()
}

func creatingSlice() {
	var slice []int = make([]int, 3) // create slice of int type with initial size & capacity of 3
	slice[0] = 1
	slice[1] = 56
	slice[2] = 74
	//adding slice[3] =5, will cause out of index error
	fmt.Println(slice)
}

func creatingSliceShortHand() () {
	slice := []int{1, 25, 3}
	fmt.Println(slice)

}

func sliceASlice() {
	slice := []int{1, 25, 6, 7, 25, 3}

	fmt.Println(slice[2:3]) //slice down a slice from 2nd index before 3rd index, last index is excluded
}

/*
We use append to increase size of a slice*/
func increasingSizeOfSlice() {
	slice := []int{1, 25}
	//slice[2]=5;//This will give us error as size of slice is already defined and we adding to it

	slice = append(slice, 5, 6, 2)// this will append 5,6,2 to the slice

	fmt.Println(slice)

	slice = append(slice, slice...)// this will append entire slice(2nd Param) to slice(first param)

	fmt.Println(slice)
}
