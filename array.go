package main

import "fmt"

func main() {
	dynamic_arr := [...]int{1, 2, 3}
	for _, value :=  range dynamic_arr {
		fmt.Println(value)
	}
}
