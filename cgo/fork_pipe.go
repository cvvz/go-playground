package main

// #include "fork_wrapper.h"
import "C"

import (
	"fmt"
	"os"
	"time"
)

func main() {
	r, w, _ := os.Pipe()

	fmt.Printf("pid %d\n", os.Getpid())
	pid := C.fork_wrapper()

	if pid == -1 {
		fmt.Println("fork failed")
		os.Exit(1)
	} else if pid > 0 {
		fmt.Printf("parent process: pid %d, child pid %d\n", os.Getpid(), pid)
		buf := make([]byte, 1024)
		timeTick := time.Tick(500 * time.Millisecond)
		for {
			select {
			case <-timeTick:
				n, _ := r.Read(buf)
				fmt.Printf("parent process get %d bytes: %s", n, buf)
			}
		}
	} else {
		fmt.Printf("child process: pid %d\n", os.Getpid())
		os.Stdout = w
		os.Stderr = w
		fmt.Println("test!")
	}
}
