package callbackAndroidGoCFunc

/*
#cgo CFLAGS: -I .
#cgo LDFLAGS: -L . -lcalculator

#include "calculator.h"

int callbackSum_cgo(int in);
int callbackCubic_cgo(int in);
*/
import "C"
import (
	"unsafe"
)

type CalculatorInterface interface {
	SumResult(int)
	SquareResult(int)
	CubicResult(int);
}

var calculatorInterface CalculatorInterface

func RegisterCalculatorInterface(c CalculatorInterface) int {
	if c == nil {
		return 1 // fail
	}
	if IsRegistered() {
		return 2 // fail
	}
	calculatorInterface = c
	return 0 // successful
}

func UnRegisterCalculatorInterface() {
	if !IsRegistered() {
		return
	}
	calculatorInterface = nil
}

func IsRegistered() bool {
	return calculatorInterface != nil
}

//export callbackSum
func callbackSum(output int) {
	if !IsRegistered() {
		return
	}

	calculatorInterface.SumResult(output)
}

func Sum(value int, value2 int) bool {
	if !IsRegistered() {
		return false
	}

	C.sum_c_func(
		(C.callbackSum_fcn)(unsafe.Pointer(C.callbackSum_cgo)),
		C.int(value),
		C.int(value2))

	return true
}

//export callbackCubic
func callbackCubic(output int) {
	if !IsRegistered() {
		return
	}

	calculatorInterface.CubicResult(output)
}

func Cubic(value int) bool {
	if !IsRegistered() {
		return false
	}

	C.cubic_c_func(
		(C.callbackCubic_fcn)(unsafe.Pointer(C.callbackCubic_cgo)),
		C.int(value))

	return true
}

func Square(value int) bool {
	if !IsRegistered() {
		return false
	}

	calculatorInterface.SquareResult(value * value)

	return true
}

func Multiply(value int, value2 int) int {
	return value * value2
}
