package singleton

import (
	"sync"
)

type Apple struct {
	name string
}

var once sync.Once
var apple *Apple

func getInstance() *Apple {
	once.Do(func() {
		apple = &Apple{}
	})
	return apple
}
