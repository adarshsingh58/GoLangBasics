package main

import "fmt"

type PrinterMethod func(string) ()

func main() {

	CreateCustomGreet("Mr")(" How are you!!!")

	GreetWithClosure("Hey Man How are you?",CreateCustomGreet("Dude"))// Passing a method
}

//Returning a method
func CreateCustomGreet(prefix string) PrinterMethod {
	return func(message string) {
		fmt.Println(prefix + " " + message)
	}
}

func GreetWithClosure(message string, do PrinterMethod){
	do(message)
}
