package main

import (
	"fmt"
	"regexp"
)

func main() {
	var s string = "a  b                c  d  e"

	reg := regexp.MustCompile(`\s+`)
	s = reg.ReplaceAllString(s, " ")
	fmt.Println(s)
}
