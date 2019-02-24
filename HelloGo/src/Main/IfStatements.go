package main

import "fmt"

//if OptionalStatement ; CONDITION {  BLOCK OF CODE }

func main() {
	GreetConditionally("Hello", false)
}
func GreetConditionally(msg string, doGreet bool) {
	if (doGreet) {
		fmt.Println(msg + " Sir how are you?")
	} else {
		fmt.Println("Sorry No Greetings!!!")
	}
}
