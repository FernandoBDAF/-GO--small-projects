package main

import (
	"fmt"
	"time"
)

func printMessage(text string, channel chan string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(800 * time.Millisecond)
	}
	channel <- "Done!"
}

func main() {
	channel := make(chan string)
	// go printMessage(("Go is fun!"))
	go printMessage("Concurrent programming is fun!", channel)
	reponse := <- channel // main routine waiting for the channel to receive a message
	fmt.Println(reponse)
	// go printMessage(("Goroutines are fun!"))
}