package main

import "fmt"

func main() {
	/**
	* !ARRAYS
	*/
	/* var fruitsArr [2] string
	fruitsArr[0] = "Orange"
	fruitsArr[1] = "Apple" */

	// short hand for declaring and assigning to an array
	//fruitsArr := [2] string {"Orange", "Apple"}
	//fmt.Println(fruitsArr)
	//fmt.Println(fruitsArr[1])

	/**
	* !SPLICES
	* number of items in array is not specified
	*/
	//var fruitArr [2]string
	//fruitArr[0] = "Orange"
	//fruitArr[1] = "Apple"

	

	//shorthand

	fruitArr := [] string {"apple","orange","mango"}

	fmt.Println(fruitArr)
	//lenght method
	fmt.Println(len(fruitArr))

	//range method
	fmt.Println(fruitArr[1:2])
	
}