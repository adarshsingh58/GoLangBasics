package main

import "fmt"

//Variadic Functions are like varargs in Java
// You can use ... to send multiple parameters with one var
//This must always be the last param
func main() {

	GreetAllPeopleVariadicFunction("Hello", "Adarsh", "Jagga", "John", "Doe")

}
func GreetAllPeopleVariadicFunction(msg string, name ...string) {
	fmt.Println(len(name))
	fmt.Println(msg + " " + name[0] + " " + name[1] + " " + name[2] + " " + name[3])
}
