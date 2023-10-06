package main

import "fmt"

type admin struct {
	adminID   string
	adminName string
	phone     string
}

func main() {
	adminList := []admin{}
	admin1 := admin{
		adminID:   "001",
		adminName: "Kenn1",
		phone:     "0010000",
	}
	admin2 := admin{
		adminID:   "002",
		adminName: "Kenn2",
		phone:     "0020000",
	}
	admin3 := admin{
		adminID:   "003",
		adminName: "Kenn3",
		phone:     "0030000",
	}

	adminList = append(adminList, admin1)
	adminList = append(adminList, admin2)
	adminList = append(adminList, admin3)

	fmt.Println(adminList)
}
