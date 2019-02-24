package main

import "fmt"

func main() {
	classicFor()
	whileUsingFor()
	classicForUsingBreak()
	classicForUsingContinue()
	forEachUsingFor()
}

func whileUsingFor() {
	i := 0
	for i < 3 {
		fmt.Println("While Loop three times")
		i++
	}
}

func classicFor() {
	for i := 0; i < 5; i++ {
		fmt.Println("For Loop five times")
	}
}

func forEachUsingFor() {
	slice := []string{"adarsh", "bob", "John"}
	//here i will be index and s will be value
	//using i, s is mandatory with range, if we do not wish to use i
	//we can use _ instead of i
	//if we use only one var like "for i := range slice " only index will be mapped and no value
	for i, s := range slice {
		fmt.Print(i)
		fmt.Println(" value=" + s)

	}
}

func classicForUsingBreak() {
	for i := 0; i < 5; i++ {
		fmt.Println("For With Break Loop 2 of 5 times")
		if i == 1 {
			break
		}
	}
}
func classicForUsingContinue() {
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("For With Continue printing only odd times")
	}
}
