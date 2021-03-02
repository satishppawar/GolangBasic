package main

import "fmt"

func main() {
	var name string
	name = "Satish"
	fmt.Println("Hello " + name + " !Welcome to Go Lang")

	//use %T to get type
	fmt.Printf("Type is %T \n", name)

	//number, err := fmt.Println("Hello", 100, true, "Go Lang is awesome") //GolangBasic/01-Go-Basic/02-DataTypes.go:13:10: err declared but not used
	number, _ := fmt.Println("Hello", 100, true, "Go Lang is awesome") //GolangBasic/01-Go-Basic/02-DataTypes.go:13:10: err declared but not used
	fmt.Println(number)

}
