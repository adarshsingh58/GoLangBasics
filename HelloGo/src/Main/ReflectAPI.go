package main

import (
	"fmt"
	"reflect"
)

var name string="Adarsh"
var address ="DSR"
var age =20

func main() {

	fmt.Println("Type of value address is", reflect.TypeOf(address))//this will tell the type of variable

}
