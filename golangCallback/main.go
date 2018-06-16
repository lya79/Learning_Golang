package main

import "fmt"

func hello(name string, callback func(string)) {
	callback("hello " + name)
}

func main() {
	fHello := func(n string) {
		fmt.Println(n) // output: hello World
	}

	hello("World", fHello)
}
