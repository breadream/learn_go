package main

func main() {
	add := func (i int, j int) int {
		return i + j
	}

	// pass `add` func
	r1 := calc(add, 10, 20)
	println(r1)

	r2 := calc(func(x int, y int) int { return x - y }, 10, 20)
	println(r2)
}

func calc(f func(int, int) int, a int, b int) int {
	result := f(a, b)
	return result
}
