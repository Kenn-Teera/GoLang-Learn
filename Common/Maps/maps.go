package main

import "fmt"

var product = make(map[string]float64)

func main() {
	fmt.Println("Product =",product)	

	//add data
	product["Macbook"] = 4000
	product["Ipand"] = 2000
	product["mama"] = 1000
	fmt.Println("Product =",product)	
	
	//delete

	delete(product,"Ipand")
	fmt.Println("Product =",product)	

	//update
	product["mama"] = 10
	fmt.Println("Product =",product)	

	value1 := product["Macbook"]
	fmt.Println("value1 =",value1)

	courseName := map[string]string{"101":"Java","102":"Python"}
	fmt.Println("courseName =",courseName)
}
