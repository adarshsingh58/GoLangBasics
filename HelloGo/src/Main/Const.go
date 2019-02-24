package main

import "fmt"

const salutationMale  ="Mr"//defining const in one line
//type automatically inferred
//can define multiple const with () in one block
const (
	salutationFemaleUnmarried="Ms"
	salutationFemaleMarried="Mrs"
	PI=3.14
	//iota represent successive untyped integer constants starting from 0
	A=iota// 3rd const inside this const block
	B=iota// 4th const inside this const block
	C=iota
)

const(
	Landlord="Mr Rag"
	Owner//when value not defined for a constant, it takes the value of previously defined const
	D=iota
	E//when value not defined for a constant, it takes the value of previously defined const
)
func main() {

	fmt.Println(PI)
	fmt.Println(salutationFemaleMarried+" Sulma")
	fmt.Println(A)
	fmt.Println(B)

	fmt.Println(Landlord)
	fmt.Println(Owner)
	fmt.Println(D)
	fmt.Println(E)
}
