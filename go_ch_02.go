package main

import "fmt"

// func main() {
// 	c := make(chan int) // unbuffered channel
// 	c <- 1 // deadlock 
// 	fmt.Println(<-c) // won't work
// }

func main() {
	c := make(chan int, 1) // buffered channel
	c <- 101
	fmt.Println(<-c)
}
