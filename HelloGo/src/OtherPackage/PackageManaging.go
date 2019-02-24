package OtherPackage

import "fmt"

//Because a go file can only run from a main method inside a 'main' package only
//we can do packaging for different modules and call them from our main module
//This module will be called from PackageMain.go file in Main folder

func GreetInSpanish(name string){
	fmt.Println("Ola "+name)
}