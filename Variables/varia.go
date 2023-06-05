package main

import "fmt"

var numberInt,numberInt2 int = 100,200
var msg string = "This is string"

func main() {
	numberfloat := 29.5
	fmt.Println(numberInt)
	fmt.Println(numberInt2)
	fmt.Println(numberfloat)
	fmt.Println(msg)

	fmt.Println(numberInt+numberInt2)
	fmt.Println(numberInt+int(numberfloat))

	fmt.Println(msg + " and this is add on")
	fmt.Println("my money =",numberInt)
	fmt.Println("this is float =>",numberfloat)
}