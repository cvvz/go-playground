package debug_test

import (
	"fmt"
	"testing"
)

func TestGODEBUG(t *testing.T) {
	// os.Setenv("GODEBUG", "backtrace=1")
	// panic("test panic")

}

// https://go.dev/doc/articles/race_detector
func TestRace(t *testing.T) {
	c := make(chan bool)
	m := make(map[string]string)
	go func() {
		m["1"] = "a" // First conflicting access.
		c <- true
	}()
	m["2"] = "b" // Second conflicting access.
	<-c
	for k, v := range m {
		fmt.Println(k, v)
	}
}
