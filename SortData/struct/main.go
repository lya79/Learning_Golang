package main

import (
	"fmt"
	"sort"
)

type Member struct {
	Id        int
	LastName  string
	FirstName string
}

type byLastFirst []Member

func (a byLastFirst) Len() int      { return len(a) }
func (a byLastFirst) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a byLastFirst) Less(i, j int) bool {
	if a[i].LastName < a[j].LastName {
		return true
	}
	if a[i].LastName > a[j].LastName {
		return false
	}
	return a[i].FirstName < a[j].FirstName
}
func main() {
	members := [6]Member{}
	members[0] = Member{LastName: "a", FirstName: "b"}
	members[1] = Member{LastName: "a3", FirstName: "b3"}
	members[2] = Member{LastName: "a2", FirstName: "b2"}
	members[3] = Member{LastName: "a3", FirstName: "b3"}
	members[4] = Member{LastName: "a4", FirstName: "b4"}
	members[5] = Member{LastName: "a1", FirstName: "b1"}

	slice := members[0:len(members)]
	sort.Sort(byLastFirst(slice))
	fmt.Println("Sorted: ", slice)
}
