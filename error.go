package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.Open("hi")
	switch err.(type) {
	default:
		println("ok")
	case error:
		log.Fatal(err.Error())
	}
	println(f.Name())
}
