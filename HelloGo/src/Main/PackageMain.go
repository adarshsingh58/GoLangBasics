package main

import "../OtherPackage"

//.. means go back twice up in folder hierarchy
//Then / as go to OtherPackage folder and all the go files inside
//it will be imported
// You must make sure that GreetInSpanish should be in FirstLetter Capital case
//to make it public.
func main() {
	OtherPackage.GreetInSpanish("Adarsh")
}
