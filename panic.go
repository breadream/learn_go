// panic() runs the all defer functions and return

package main

import "os"

func main() {
	openFile("hi")

	// it won't run
	println("Done")
}

func openFile(fn string) {
	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}
	defer f.Close()
}
