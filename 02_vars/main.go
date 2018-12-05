package main

import "fmt"

func main() {
	// using vars
	// var name = "Johnson"
	var age = 35

	//using const
	const isCool = false
	
	// shorthand version, used only within a function, and not outside
	// name := "Johnson"
	size := 2.3
	// email := "ogwurujohnson"

	//we can also do the following
	name, email := "Johnson","ogwurujohnson@gmail.com"

	fmt.Println(name, age, isCool, size, email)
	fmt.Printf("%T\n", size)
}

