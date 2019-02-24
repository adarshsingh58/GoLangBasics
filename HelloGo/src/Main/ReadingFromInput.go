package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	reader:=bufio.NewReader(os.Stdin)
	fmt.Println("Enter Text")
	text,_:=reader.ReadString('\n')
	fmt.Println(text)

}
