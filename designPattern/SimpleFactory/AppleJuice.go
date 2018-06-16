package SimpleFactory

type AppleJuice struct {
	name string
}

func (a *AppleJuice) getName() string {

	a.name = "Apple Juice"

	return a.name
}
