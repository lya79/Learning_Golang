package main

import (
	"fmt"
	"reflect"
)

type human struct {
	name string
}

type student struct {
	h  human
	id int
}

func main() {
	var a1 student
	a1.h = human{name: "HUMAN"}
	a1.id = 111

	var a2 student
	a2.h = human{name: "HUMAN"}
	a2.id = 111

	var a3 student
	a3.h = human{name: "HUMAN"}
	a3.id = 1112

	fmt.Println(a1, "==", a2, reflect.DeepEqual(a1, a2))
	fmt.Println(a1, "==", a3, reflect.DeepEqual(a1, a3))

	b1 := []int{1, 2}
	b2 := []int{1, 2}
	if reflect.DeepEqual(b1, b2) {
		fmt.Println(b1, "==", b2)
	}

	c1 := map[string]int{"a": 1, "b": 2}
	c2 := map[string]int{"a": 1, "b": 2}
	if reflect.DeepEqual(c1, c2) {
		fmt.Println(c1, "==", c2)
	}
}
