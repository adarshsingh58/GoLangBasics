package main

import (
	"fmt"
)

func main() {
	printSeq := false
	if (printSeq) {
		printAA()
		printBB()
	} else {
		buffer := make(chan bool)// we created a channel/pipe/buffer of type boolean
		go func() {// new thread spawned, we used anonymous func here, so that buffer<-true can be a part of the thread too
			printAA()
			buffer <- true//buffer will be pushed in with value true once this thread id completed
			//only one buffer <- true can be used here, since it is a channel and not buffered channel
			//if we do, buffer <- true buffer <- true two times consecutively then our code will be blocked,
			//because in normal channel, one 1 data is added to the channel, it must be removed before another data
			//can be inserted or added.
			buffer <- true// control blocked here, by the time data from buffer is read, main thead execution might be over
			fmt.Println("Done") // hence, This will not be printed
		}()//anon func executed
		printBB()//This func is called in main thread
		<-buffer// we read a value from the channel and only when value is available in channel is this line processed
				//else, control will wait on this line for some value to be pushed in the channel
		}
}

func printAA() {
	for i := 1; i <= 5; i++ {
		fmt.Println("A", i)
	}
}

func printBB() {
	for i := 1; i <= 5; i++ {
		fmt.Println("B", i)
	}
}
