package testmain

// $ go test
// 2023/03/17 00:19:15 Do stuff BEFORE the tests!
// 2023/03/17 00:19:15 TestA running
// 2023/03/17 00:19:15 TestB running
// PASS
// 2023/03/17 00:19:15 Do stuff AFTER the tests!
// ok      example/testmain        0.002s

import (
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	log.Println("Do stuff BEFORE the tests!")
	exitVal := m.Run()
	log.Println("Do stuff AFTER the tests!")

	os.Exit(exitVal)
}

func TestA(t *testing.T) {
	log.Println("TestA running")
}

func TestB(t *testing.T) {
	log.Println("TestB running")
}
