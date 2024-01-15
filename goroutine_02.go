// go routine can be used on anonymous function

package main
import (
	"fmt"
	"sync"
)

func main() {
	// create WaitGroup, wait for 2 goroutines
	var wait sync.WaitGroup
	wait.Add(2)

	// goruotine using anonymous func
	go func() {
		defer wait.Done()
		fmt.Println("Hello")
	}()

	// passing an argument to anonymous func
	go func(msg string) {
		defer wait.Done()
		fmt.Println(msg)
	}("hi")

	wait.Wait() // Wait for all goruotines finish
}
