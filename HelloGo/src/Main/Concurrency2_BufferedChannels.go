package main

import "fmt"

func main() {
	printSeq := false
	if (printSeq) {
		printAAA()
		printBBB()
	} else {
		buffer := make(chan bool, 2)// channel with buffer size 2 defined
		go func() {
			printAAA()
			buffer <- true
			buffer <- true// hence we were able to add 2 data in channel, hence next line is not blocked
			fmt.Println("Done")// we will see Done printed on console
		}()
		printBBB()
		<-buffer//reading from channel
	}
}

func printAAA() {
	for i := 1; i <= 5; i++ {
		fmt.Println("A", i)
	}
}

func printBBB() {
	for i := 1; i <= 5; i++ {
		fmt.Println("B", i)
	}
}
