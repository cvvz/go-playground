package generic

import "testing"

type S struct{}

func (s *S) example(t int) {
	println(t)
}

func (s *S) example2(t int) {
	println(t)
}

func (s *S) example3(t float64) {
	println(t)
}

type S2 struct{}

func (s *S2) example(t string) {
	println(t)
}
func (s *S2) example2(t string) {
	println(t)
}
func (s *S2) example3(t float64) {
	println(t)
}
func TestGenericInterface(t *testing.T) {
	// 在接口实例化的时候，就把T的具体类型(int还是string)定下来了

	// genericInterface[int]
	var _ genericInterface[int] = &S{}

	// genericInterface[string]
	var _ genericInterface[string] = &S2{}

}
