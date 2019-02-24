package main

import "fmt"

/*
>No default fall through i.e. once a case is matched the call will come out
of the entire switch block, so no need to put breaks in every case.
If we want a fall through, we can manually do that

>cases can be expression here

>We can switch based on type of input to switch

>we can multiple cases in single case block
*/
func main() {
	fmt.Println("Hello " + getPrefix("Male") + "Adarsh")
	fmt.Println("Hello " + getPrefix("Femail") + "Jinga")
	printPrefixWithFallthrough("Male")
	printPrefixBasedOnTypeOnInput("Male")
	printPrefixBasedOnTypeOnInput(6)

}
func getPrefix(gender string) (prefix string) {

	switch gender {
	case "Male":
		prefix = "Mr."
	case "Female", "Femail": //we can multiple cases in single case block
		prefix = "Mrs."
	default:
		prefix = "UNKNOWN"
	}
	return prefix
}

//Here because we have used fallthrough on Male case, the control
//will fallthrough the next case as well which is the female case
func printPrefixWithFallthrough(gender string) {
	switch gender {
	case "Male":
		fmt.Println("Hello Male")
		fallthrough
	case "Female":
		fmt.Println("Hello FeMale")
	default:
		fmt.Println("Hello Mail")
	}
}

//x can be of anytype. Equi to Object in Java
func printPrefixBasedOnTypeOnInput(x interface{}) {
	switch x.(type) {// this will get the type of x
	case int:
		fmt.Println("Integer Type");
	case string:
		fmt.Println("String Type")
	default:
		fmt.Println("No type")
	}
}
