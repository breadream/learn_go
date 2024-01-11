// struct only has fields
// method specifies its struct to support

package main

type Rect struct {
	width, height int
}

func (r Rect) area() int {
	return r.width * r.height
}

// value receiver vs pointer receiver
func (r *Rect) doubleArea() int {
	r.width++
	return r.width * r.height
}

func main() {
	rect := Rect{10, 20}
	area := rect.area() // calling a method
	println(area)

	secondArea := rect.doubleArea()
	println(secondArea)
}
