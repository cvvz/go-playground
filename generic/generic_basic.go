package generic

// 类型约束(Type constraint)
type Number interface {
	int64 | float64
}

func SumIntAndFloat[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func Sum[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
