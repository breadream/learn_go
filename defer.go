// `defer` works as finally block

package main

import "os"

func main() {
	f, err := os.Open("hi")
	if err != nil {
		panic(err)
	}

	// run file close at the end
	defer f.Close()

	bytes := make([]byte, 1024)
	f.Read(bytes)
	println(len(bytes))
}
