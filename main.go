package main

import (
	"fmt"
)

func main(){
	fmt.Println("Hello World")

	//var myName string = "Aman Jain"
	myName := "aman jain"
	myInt := 10
	myFloat := 10.10
	fmt.Println(myName)
	fmt.Printf("My name is %s and my int is %d and my lfoat is %f\n", myName,myInt,myFloat)

	// go goes down to zero 

	var myFrndName string
	var myAge int
	var myGrade bool

	myFrndName = "Prime"
	fmt.Printf("my frnd name is %s and his age is %d and he is passed %t\n", myFrndName, myAge, myGrade) // output _ 0 false

	// some ds might go to nil value if not initialized
}