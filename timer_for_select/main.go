package main

import (
	"fmt"
	"time"

	mount "k8s.io/mount-utils"
)

func main() {
	err := waitForMount("/tmp", 10*time.Millisecond, 2*time.Second)
	if err != nil {
		fmt.Println(err)
	}
}

func waitForMount(path string, intervel, timeout time.Duration) error {
	timeAfter := time.After(timeout)
	timeTick := time.Tick(intervel)

	for {
		select {
		case <-timeTick:
			notMount, err := mount.New("").IsLikelyNotMountPoint(path)
			if err != nil {
				return err
			}
			if !notMount {
				fmt.Printf("blobfuse mount at %s success\n", path)
				return nil
			}
			fmt.Printf("blobfuse still NOT mount at %s\n", path)
		// Because time.After is a function, so on every iteration it returns a new channel. If you want this channel to be the same for all iterations, you should save it before the loop
		// https://stackoverflow.com/questions/35036653/why-doesnt-this-golang-code-to-select-among-multiple-time-after-channels-work
		// ☹️case <-time.After(timeout):
		case <-timeAfter:
			return fmt.Errorf("timeout waiting for mount %s", path)
		}
	}
}
