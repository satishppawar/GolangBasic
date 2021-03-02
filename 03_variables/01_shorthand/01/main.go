package main

import "fmt"

func main() {

	// := short declaration operator
	a := 10 //Delcare and assign
	//a := 200 //Not allowed  go:8:4: no new variables on left side of :=
	a = 100 // re-assign value
	b := "golang"
	c := 4.17
	d := true
	e := "Hello"
	f := `Do you like my hat?`
	g := 'M'

	//z := a + b + c + d + e + f + g

	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)
	fmt.Printf("%v \n", e)
	fmt.Printf("%v \n", f)
	fmt.Printf("%v \n", g)
	//fmt.Printf("%v \n", z)
}
