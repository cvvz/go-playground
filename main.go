package main

func main() {
	var fns []func()

	for i := 0; i < 5; i++ {
		// i := i
		fns = append(fns, func() {
			a := &i
			println(a)
		})
	}

	for _, fn := range fns {
		fn()
	}
}
