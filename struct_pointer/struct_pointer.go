package main

import (
	"fmt"
	"time"
)

type ininer struct {
	date    time.Time
	mapping map[string]string
}

type inner struct {
	date time.Time
	ininer
}

type client struct {
	i  inner
	ip *inner
}

func (c client) fun() {
	c.ip.date = time.Now()
	c.ip.mapping["test"] = "test"
}

func main() {
	c := client{
		i: inner{},
		ip: &inner{
			ininer: ininer{
				mapping: make(map[string]string),
			},
		},
	}
	fmt.Println(c.ip.date)
	fmt.Println(c.ip.mapping["test"])
	c.fun()
	fmt.Println(c.ip.date)
	fmt.Println(c.ip.mapping["test"])
}
