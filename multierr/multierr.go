package main

import (
	"errors"
	"fmt"

	"go.uber.org/multierr"
)

type CustomErr struct{}

func (e *CustomErr) Error() string { return "custom error" }

var Err1 = &CustomErr{}
var Err2 = errors.New("err2")
var Err3 = errors.New("err3")

func main() {
	err := multierr.Combine(Err1)
	err = multierr.Combine(err, Err2)
	err = multierr.Combine(err, Err3)
	var customErr *CustomErr
	if errors.As(err, &customErr) && errors.Is(err, Err2) && errors.Is(err, Err3) {
		fmt.Printf("wow!\n%s\n%s\n", err.Error(), customErr.Error())
	}
}
