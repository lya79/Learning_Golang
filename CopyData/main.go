package main

import (
	"fmt"
	"reflect"
)

type human struct {
	name string
}

type student struct {
	h  human
	id int
}

func main() {
	copyStruct()
	fmt.Println("-------------------")
	copyArray()
	fmt.Println("-------------------")
	copySlice()
	fmt.Println("-------------------")
	copyMap()
}

func copyMap() {
	a := make(map[int]map[int]int)
	a[1] = make(map[int]int)
	a[1][11] = 111

	a[2] = make(map[int]int)
	a[2][22] = 222

	b := clone(a)
	fmt.Println(a, b) // map[1:map[11:111] 2:map[22:222]] map[1:map[11:111] 2:map[22:222]]
}

func copySlice() {
	arr := [][]int{{1, 2, 3}, {2, 3, 4}, {3, 4, 5}}
	arr2 := clone(arr)
	fmt.Println(arr, arr2) // [[1 2 3] [2 3 4] [3 4 5]] [[1 2 3] [2 3 4] [3 4 5]]
}

func copyArray() {
	arr := [3][3]int{{1, 2, 3}, {2, 3, 4}, {3, 4, 5}}
	arr2 := clone(arr)
	fmt.Println(arr, arr2) // [[1 2 3] [2 3 4] [3 4 5]] [[1 2 3] [2 3 4] [3 4 5]]
}

func copyStruct() {
	var a student
	a.h = human{name: "HUMAN"}
	a.id = 111

	b := clone(a)

	a2 := &a

	fmt.Println(a, a2, b)    // {{HUMAN} 111} &{{HUMAN} 111} {{HUMAN} 111}
	fmt.Println(&a, &a2, &b) // &{{HUMAN} 111} 0xc042078018 0xc04204c1c0
}
func clone(i interface{}) interface{} {
	return reflect.Indirect(reflect.ValueOf(i)).Interface()
}
