package main

func main() {
	ch := make(chan int)

	go func() {
		ch <- 123
	}()

	i := <- ch // get 123 from channel
	println(i)
}
