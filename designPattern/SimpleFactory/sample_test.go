package SimpleFactory

import "testing"

func TestApple(t *testing.T) {
	Juice := newFruitJuice(Apple)
	s := Juice.getName()
	if s != "Apple Juice" {
		t.Fatal("newFruitJuice(Apple),", s)
	}
}

func TestPear(t *testing.T) {
	Juice := newFruitJuice(Pear)
	s := Juice.getName()
	if s != "Pear Juice" {
		t.Fatal("newFruitJuice(Pear),", s)
	}
}
