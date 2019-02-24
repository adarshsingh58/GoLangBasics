package main

import "fmt"

func main() {
	emptyInterface("adarsh")
	emptyInterface(5)
}
/*
An empty interface in Go is an interfac which is implemented by all the types.
Hence this can be used as a generic type
*/
func emptyInterface(x interface{}) {
	switch x.(type) {
	case int:
		{
			fmt.Println("Integer Type")
		}
	case float32:
		{
			fmt.Println("Float Type")
		}
	case string:
		{
			fmt.Println("String Type")
		}

	}
}
