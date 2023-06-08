package main

import "fmt"

func main() {
	var count = 5

	i := 0
	for i < 10 {
		fmt.Println("i =",i)
		i++
	}

	for j := 0; j < count; j++ {
		fmt.Println("j = ",j)
	}

	var input string
	for {
		fmt.Scanf("%s", &input)
		fmt.Println("input = ",input)
		if input == "exit" {
			break
		}
	}
}