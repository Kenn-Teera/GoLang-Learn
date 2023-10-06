package main

import "fmt"

func main() {
	var Name []string
	Name = []string{"Python","Java"}
	fmt.Println(Name)
	Name = append(Name, "C")
	fmt.Println(Name)
	Name = append(Name, "C#","HTML","JSON","CSS","Javascript")
	fmt.Println(Name)
	
	CName := Name[3:7]
	fmt.Println(CName)
}