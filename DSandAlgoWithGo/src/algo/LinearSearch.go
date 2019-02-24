package main

import (
	"fmt"
)

func main() {
	/*fmt.Println("Program Searching data in list Sequentially!!!")
	fmt.Println("Enter number of input elements")

	var i int
	_, err := fmt.Scanf("%d", &i)
	if err==nil {
		fmt.Printf("Enter %d elements",i)
	}else{
		fmt.Println(err)
	}*/

	arr := []int {1,2,3,4,5,6,112,32}
	op:=linearSearch(arr,3)
	fmt.Println(op)
}

func linearSearch(arr []int, num int) (index int){
	for i,v:= range arr{
		if v == num{
			index = i
			break
		}
	}
	return
}
