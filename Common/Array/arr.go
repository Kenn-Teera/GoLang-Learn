package main

import "fmt"

var product [4]string


func main() {
	product[0] = "Ley"
	product[1] = "Coke"
	product[2] = "Beer"
	product[3] = "Mama"

	price := [4]float64{10,20,40,5}

	fmt.Println(product)
	fmt.Println(price)
}