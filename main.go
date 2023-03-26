package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Println(filepath.Dir(strings.TrimSuffix("/", "/")))
}
