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
	for {
		select {
		case <-time.After(intervel):
			notMount, err := mount.New("").IsLikelyNotMountPoint(path)
			if err != nil {
				return err
			}
			if !notMount {
				fmt.Printf("blobfuse mount at %s success\n", path)
				return nil
			}
		case <-time.After(timeout):
			return fmt.Errorf("timeout waiting for mount %s", path)
		}
	}
}
