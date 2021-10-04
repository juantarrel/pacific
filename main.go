package main

import "fmt"
type fp struct {
	fs []func()
}

func(f *fp) run() {
	for _, fn := range f.fs {
		fn()
	}
}

func Compose(functions ...func()) *fp {
	compose := &fp{
		fs: append(functions),
	}
	compose.run()
	return compose
}

func print1()  {
	fmt.Println("hola1")
}
func print2()  {
	fmt.Println("hola2")

}

func main() {
	Compose(
		print1,
		print2,
	)
}
