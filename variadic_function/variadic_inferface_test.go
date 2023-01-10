package variadicfunction

import "testing"

func Test_funcWithVariadicInterface(t *testing.T) {
	funcWithVariadicInterface(func() { println("this is the second function") }, 1, 2, 3, "test", func() { println("this is the second function") })
}
