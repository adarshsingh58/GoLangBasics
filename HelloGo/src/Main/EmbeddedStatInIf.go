package main

import "fmt"
//if OptionalStatement ; CONDITION {  BLOCK OF CODE }
// OptionalStatement here is assignment of a var which exist only in the scope of IF and ELSE block
func main() {
	isExist := false;

	if name:="Adarsh"; (isExist) {
		fmt.Println(name + " exist")
	} else {
		fmt.Println(name+" does not one exist")
	}

	//fmt.Println(name + " exist") Here it wont work
}
