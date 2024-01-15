package main

import (
	"errors"
	"fmt"
	"testing"
)

type MyError struct {
	msg string
}

func (e *MyError) Error() string {
	return e.msg
}

func TestErrorsAsIs(t *testing.T) {
	err := foo()
	var myerr2 *MyError

	if errors.Is(err, myerr2) {
		fmt.Printf("Is myerr2, %s\n", err)
	}

	if errors.Is(err, &MyError{}) {
		fmt.Printf("Is MyError, %s\n", err)
	}

	if errors.As(err, &myerr2) {
		fmt.Printf("As MyError, %s\n", err)
	}
}

func foo() error {
	err := &MyError{
		msg: "my error",
	}
	return fmt.Errorf("error: %w", err)
}
