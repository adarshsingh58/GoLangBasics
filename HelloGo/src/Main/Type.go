package main


import "fmt"

type Salutation struct {
	name     string
	greeting string
}

func main() {
	var s1 Salutation = Salutation{"Joe", "Hello Joe"}
	var s2 = Salutation{"Molly", "Hello Molly"}
	s3 := Salutation{"Dan", "Hello Dan"}
	var s4 Salutation = Salutation{greeting: "Mahlao", name: "Kane"}
	var s5 Salutation = Salutation{}
	s5.name = "Adarsh"
	s5.greeting = "Namaste"

	fmt.Println(s1)
	fmt.Println(s1.name, s1.greeting)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println(s5)
}

