package interfacenil_test

import (
	"testing"
)

type Interface interface {
	Do()
}

type Struct struct{}

func (s *Struct) Do() {}

var _ Interface = (*Struct)(nil)

func TestInterfaceNil(t *testing.T) {
	var i Interface // 接口本身为nil
	// i.Do()       // panic: runtime error: invalid memory address or nil pointer dereference

	var ps *Struct // 接口的实例为nil
	i = ps
	i.Do() // 不会panic，因为虽然ps为nil，但是ps的类型为*Struct，所以i的类型为*Struct，所以i不为nil，所以i.Do()不会panic
}
