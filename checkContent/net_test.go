package checkContent

import (
	"testing"
)

func TestCheckIPv4Addr(t *testing.T) {
	var err error
	var addr string

	addr = "123.123.123.123:456"
	err = CheckIPv4Addr(addr)
	if err == nil {
		t.Log("Pass" + addr)
	} else {
		t.Error("Error" + addr)
	}

	addr = "123.123.123.123:-1"
	err = CheckIPv4Addr(addr)
	if err == nil {
		t.Error("Error" + addr)
	} else {
		t.Log("Pass" + addr)
	}

	addr = "123.123.123:456"
	err = CheckIPv4Addr(addr)
	if err == nil {
		t.Error("Error" + addr)
	} else {
		t.Log("Pass" + addr)
	}

	addr = "123.123.123,456"
	err = CheckIPv4Addr(addr)
	if err == nil {
		t.Error("Error" + addr)
	} else {
		t.Log("Pass" + addr)
	}

	addr = "123.123.123:"
	err = CheckIPv4Addr(addr)
	if err == nil {
		t.Error("Error" + addr)
	} else {
		t.Log("Pass" + addr)
	}
}
