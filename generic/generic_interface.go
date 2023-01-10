package generic

// 带泛型的Basic Interface
type genericInterface[T int | string] interface {
	example(t T)
	example2(t T)
	example3(t float64)
}

// General Interface
type Uint interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

// 错误。Uint是一般接口，只能用于类型约束，不得用于变量定义
var uintInf Uint 
