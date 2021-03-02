package main //Entry point of every program is main package

import "fmt"

func main() { //mandatory
	foo()

	fmt.Println("Welcome to Go Lang")
	fmt.Println("Welcome to the GO lang proamming ")
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println("Value of i is", i)
		}
	}
	bar()
}

//function2

func foo() {
	fmt.Println("Executing foo...")
}

func bar() {
	fmt.Println("Program execution is completed...")
}
