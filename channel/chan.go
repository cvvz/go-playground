package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan interface{})
	go func() {
		for {
			<-c
			fmt.Println("Unblocked1.")
		}
	}()
	close(c)
	// c = make(chan interface{})
	go func() {
		for {
			<-c
			fmt.Println("Unblocked2.")
		}
	}()
	// close(c)
	// c = make(chan interface{})
	fmt.Println("Done.")
	time.Sleep(1 * time.Second)
}
