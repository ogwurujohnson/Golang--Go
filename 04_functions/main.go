package main

import "fmt"

func greeting(name string) string {
	return "Hello " + name
}

func sum(num1, num2 int) int {
	return num1 + num2
}

func diff(num1, num2 float32) float32 {
	return num1 - num2
}

func multiply(num1, num2 int) int {
	return num1 * num2
}

func main() {
	fmt.Println(greeting("Johnson"))
	fmt.Println(sum(2,4))
	fmt.Println(diff(5.4, 3.2))
	fmt.Println(multiply(2,2))
}

