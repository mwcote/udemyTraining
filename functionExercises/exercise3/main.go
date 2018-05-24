package main

import "fmt"

func main() {
	fmt.Println(findMax(1, 5, 7, 12, 14, 12, 11, 19, 28, 13, 27))
}

func findMax(lst ...int) int {
	var maximum int = lst[0]
	for _, v := range lst {
		if v > maximum {
			maximum = v
		}
	}
	return maximum
}
