package main

import "fmt"

func main() {
	foo(1, 2, 3)
	foo(1, 2)
	aSlice := []int{1, 2, 3, 4}
	foo(aSlice...)
	foo()
}

func foo(nums ...int) {
	fmt.Println("FOOOOOOOOOOO")
}
