package main
import "fmt"

func main() {
//	sliceA := make([]int, 0, 3)
//
//	for i := 1; i <= 15; i++ {
//		sliceA = append(sliceA, i)
//		fmt.Println(len(sliceA), cap(sliceA))
//	}
//	fmt.Println(sliceA)
	
// 	sliceA := []int{1, 2, 3}
// 	sliceB := []int{4, 5, 6}
// 
// 	sliceA = append(sliceA, sliceB...)
// 
// 	fmt.Println(sliceA)

	source := []int{0, 1, 2}
	target := make([]int, len(source), cap(source)*2)
	copy(target, source)
	fmt.Println(target)
	println(len(target), cap(target))
}
