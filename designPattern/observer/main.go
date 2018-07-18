package main

import (
	"fmt"
)

func main() {
	c := &client{}
	c.name = "client 1"

	c2 := &client{}
	c2.name = "client 2"

	s := server{}
	s.observerMap = make(map[observer]bool)
	s.addListener(c)
	s.addListener(c2)

	s.notify()
}

//---------------------------------

type observer interface {
	update()
}

//---------------------------------

type subject interface {
	addListener(observer)
	removeListener(observer)
	notify()
}

//---------------------------------

type client struct {
	name string
}

func (c *client) update() {
	fmt.Println("update, name:", c.name)
}

//---------------------------------

type server struct {
	observerMap map[observer]bool
}

func (s *server) addListener(ob observer) {
	s.observerMap[ob] = true
}

func (s *server) removeListener(ob observer) {
	delete(s.observerMap, ob)
}

func (s *server) notify() {
	for ob, _ := range s.observerMap {
		ob.update()
	}
}
