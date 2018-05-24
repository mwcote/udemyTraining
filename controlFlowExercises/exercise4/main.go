package main

import "fmt"

func main() {
	var bigNum int
	var smallNum int

	fmt.Printf("Enter a Small Number: ")
	fmt.Scan(&smallNum)
	fmt.Printf("Enter a Big Number: ")
	fmt.Scan(&bigNum)

	fmt.Println("Big Num", bigNum, "divided by", smallNum, "has remainder", (bigNum % smallNum))
}
