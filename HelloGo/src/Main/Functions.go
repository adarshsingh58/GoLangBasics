package main

import "fmt"

type Salutations struct {
	name     string
	greeting string
}

func main() {
	name := "Adarsh"
	greeting := "Namaste"
	GreetPerson(name, greeting)

	sal := Salutations{"John", "Hello"}
	GreetWithSalutation(sal)

	fmt.Println(GreetWithReturn("John", "Hello"));

	//Multiple return
	mainMsg,alternateMsg:=GreetWithMultipleReturn("Adarsh","Namaste");

	fmt.Println(mainMsg,alternateMsg)

	//Multiple return but using only selected return
	//If we have mainMsg1 and do not use it, then that will be a compile time error,
	//But 2 return type must be there, so we use _ to mark those return types which wont be used
	_,alternateMsg1:=GreetWithMultipleReturn("Jagga","Jai Ho ");

	fmt.Println(alternateMsg1)

	//Named Return
	mainMsg2,alternateMsg2:=GreetWithMultipleNamedReturn("Jagga","Jai Ho ");
	fmt.Println(mainMsg2,alternateMsg2)
}
func GreetWithMultipleNamedReturn(name , message string) (mainMsg, altMsg string) {
	mainMsg=message + " " + name
	altMsg=" How are you doing? This is multiple Named return."
	return
}

func GreetWithReturn(name, message string) string {
	return message + " " + name + ". This is from return method"
}

func GreetPerson(name string, message string) {
	fmt.Println(message + " " + name)
}

func GreetWithSalutation(salutation Salutations) {
	fmt.Println(salutation.greeting + " " + salutation.name)
}

func GreetWithMultipleReturn(name, message string) (string, string) {
	return message + " " + name, " How are you doing? This is multiple return."
}
