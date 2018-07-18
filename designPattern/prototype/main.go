package main

import (
	"fmt"
)

func main() {
	a := apple{}
	a.color = 1
	a.price = 3
	fmt.Println("a:", a)

	a2 := a.clone()
	a2.price = 5
	fmt.Println("a:", a)
	fmt.Println("a2:", a2)
}

type apple struct {
	color int
	price int
}

func (a *apple) clone() *apple {
	c := &apple{}
	c.color = a.color
	c.price = a.price
	return c
}
