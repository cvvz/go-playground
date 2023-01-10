package variadicfunction

func funcWithVariadicInterface(interfaces ...interface{}) {
	for _, i := range interfaces {
		switch v := i.(type) {
		case int:
			println(v)
		case string:
			println(v)
		case func():
			v()
		}
	}
}
