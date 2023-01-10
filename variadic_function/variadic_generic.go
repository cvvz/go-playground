package variadicfunction

type types interface {
	int | string | func()
}

func funcWithVariadicGeneric[T types](t ...T) {
	for _, i := range t {
		// 泛型不支持 动态判断变量的类型
		// https://segmentfault.com/a/1190000041634906#:~:text=%E7%9A%84%E9%98%9F%E5%88%97%0A//%20......-,4.2%20%E5%8A%A8%E6%80%81%E5%88%A4%E6%96%AD%E5%8F%98%E9%87%8F%E7%9A%84%E7%B1%BB%E5%9E%8B,-%E4%BD%BF%E7%94%A8%E6%8E%A5%E5%8F%A3%E7%9A%84
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
