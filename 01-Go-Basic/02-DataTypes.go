package main

import "fmt"

func main() {
	var name string
	name = "Satish"
	fmt.Println("Hello " + name + " !Welcome to Go Lang")

	//use %T to get type
	fmt.Printf("Type is %T", name)
}
