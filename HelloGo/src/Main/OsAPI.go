package main

import (
	"fmt"
	"os"
)

func main() {
	for a, b := range os.Environ() {
		fmt.Println(a, b)
	}
}

func print(name string) ( error){
	fmt.Println(name)
	return nil;
}
