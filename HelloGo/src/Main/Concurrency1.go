package main

import (
	"fmt"
	"time"
)

func main() {
	printSeq := false
	if (printSeq) {
		//This is sequential processing, coz both the task are happening are in single thread
		printA()
		printB()
	} else {
		//Using the keyword 'go' we make the particular call to function to spawn in a new thread
		//Here processing is happening, in 2 thread printB() in main thread, while printA() in a newly
		//spawned thread. We use sleep, coz B prints and function stops before A gets the chance to execute.
		go printA()
		printB()
		time.Sleep(10 * time.Millisecond)
	}
}

func printA() {
	for i := 1; i <= 5; i++ {
		fmt.Println("A", i)
	}
}

func printB() {
	for i := 1; i <= 5; i++ {
		fmt.Println("B", i)
	}
}
