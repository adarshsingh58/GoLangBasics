package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	pa:=Person{"asa",3}
	Person.printName(pa,"Adarsh")
}

func (p Person) printName(name string) {
	fmt.Println(name)
}
