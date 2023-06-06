package main

import "fmt"

type CloudProvider string

const (
	CloudProviderAzure CloudProvider = "Azure"
)

func main() {
	var c CloudProvider
	c = "hehehehe"
	fmt.Print(c == CloudProviderAzure)
}
