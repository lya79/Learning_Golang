package SimpleFactory

type PearJuice struct {
	name string
}

func (a *PearJuice) getName() string {
	a.name = "Pear Juice"
	return a.name
}
