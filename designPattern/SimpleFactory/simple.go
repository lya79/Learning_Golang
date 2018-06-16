package SimpleFactory

type Fruit int

const (
	Apple Fruit = iota
	Pear  Fruit = iota
)

type IJuice interface {
	getName() string
}

func newFruitJuice(fruit Fruit) IJuice {
	switch fruit {
	case Apple:
		return &AppleJuice{}
	case Pear:
		return &PearJuice{}
	default:
		return nil
	}
	return nil
}
