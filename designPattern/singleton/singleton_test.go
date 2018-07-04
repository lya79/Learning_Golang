package singleton

import (
	"testing"
)

func Test_getInstance(t *testing.T) {
	a := getInstance()
	b := getInstance()
	if a != b {
		t.Fail()
	}
	
	c := &Apple{}
	if a == c  ||  b == c{
		t.Fail()
	}
}
