package main

import (
	"fmt"
)

func main() {
	build := dateBuilder()
	data := build.setYear(2018).setMonthOfYear(7).setDayOfMonth(17).build()
	fmt.Println(data)

	build = dateBuilder()
	data = build.setMonthOfYear(7).setDayOfMonth(17).build()
	fmt.Println(data)
}

func dateBuilder() date {
	return date{}
}

type date struct {
	year        int
	monthOfYear int
	dayOfMonth  int
}

func (d *date) setYear(value int) *date {
	d.year = value
	return d
}

func (d *date) setMonthOfYear(value int) *date {
	d.monthOfYear = value
	return d
}

func (d *date) setDayOfMonth(value int) *date {
	d.dayOfMonth = value
	return d
}

func (d *date) build() *date {
	clone := &date{}
	clone.year = d.year
	clone.monthOfYear = d.monthOfYear
	clone.dayOfMonth = d.dayOfMonth
	return clone
}
