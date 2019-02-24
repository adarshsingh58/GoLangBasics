package main

import "fmt"

type Printer func(string) ()

func main() {
	Greet("Adarsh", Print)
	GreetUsingType("Adarsh", Print)
}

//Passing function as a param type
func Greet(name string, output func(string)) {
	output("Hello " + name)
}

//Passing function as a param type. Here instead of declaring func in param we used type
func GreetUsingType(name string, output Printer) {
	output("Hello " + name + ". This is using type")
}

func Print(s string) {
	fmt.Print(s)
}
func PrintLine(s string) {
	fmt.Println(s)
}
