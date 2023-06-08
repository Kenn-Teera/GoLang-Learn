package main

import "fmt"

func hello() {
	fmt.Println("Hello world!")
}
func plus(value1 int ,value2 int){
	result := value1 + value2
	fmt.Println("result =",result)
}
func minus(value1 int ,value2 int) int {
	return value1 - value2
}
func expo(value1 , value2 , value3 int) int {
	return value1 * value2 * value3
}
func main() {
	hello()
	plus(2,3)
	result := minus(5,4)
	fmt.Println("serult =",result)

	result2 := expo(1,2,3)
	fmt.Println("Expo =",result2)
}