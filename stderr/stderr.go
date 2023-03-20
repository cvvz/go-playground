package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	f, _ := os.OpenFile("test.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND|os.O_SYNC, 0755)
	os.Stdout = f
	os.Stderr = f

	os.Remove("test.txt")

	_, err := fmt.Printf("test!\n")
	if err != nil {
		os.Exit(1)
	}
	// restore stdout
	os.Stdout = os.NewFile(uintptr(syscall.Stdout), "/dev/stdout")
	fmt.Printf("test2!\n")
}
