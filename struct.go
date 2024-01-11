package main

import "fmt"

type dict struct {
	data map[int]string
}

// define constructor function
func newDict() *dict {
	d := dict{}
	d.data = map[int]string{}
	return &d // pass pointer
}

func main() {
	d := newDict()
	d.data[1] = "A"
	fmt.Println(d)
}
