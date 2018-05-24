package main

import "fmt"

func main(){
	var yourName string;
	fmt.Printf("Enter Your Name: ")
	fmt.Scan(&yourName);

	fmt.Println("Hello", yourName);
}