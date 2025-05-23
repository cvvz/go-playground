package main

import (
	"fmt"
)

func main() {
	test2()
}

func test1() (a int, err error) {
	return 1, fmt.Errorf("error")
}

func test2() (a int, err error) {
	defer func() {
		fmt.Println(err)
	}()

	c, err := test1()
	fmt.Println(c)
	return c, err
}
