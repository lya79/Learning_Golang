package main

import (
	"fmt"
)

func hello(name string, callback func(string)) {
	callback("hello " + name)
}

func main() {
	fHello := func(n string) {
		fmt.Println(n) // output: hello World
	}

	hello("World", fHello)
}

//-----------------------------------------------------------------------------------------------

// type D interface {
// 	B(s string)
// }

// type S struct {
// }

// func (s S) B(str string) {
// 	fmt.Println(str)
// }

// func main() {
// 	s := S{}
// 	hello(s)
// }

// func hello(d D) {
// 	d.B("hello world")
// }
