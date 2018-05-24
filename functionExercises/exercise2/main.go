package main

import "fmt"

func main() {
	x := func(num int) (int, bool) {
		return num / 2, (num%2 == 0)
	}
	fmt.Println(x(1))
	fmt.Println(x(2))
}
