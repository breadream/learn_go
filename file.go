package main

import "io/ioutil"

func main() {
	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("output.txt", bytes, 0)
	if err != nil {
		panic(err)
	}
}
