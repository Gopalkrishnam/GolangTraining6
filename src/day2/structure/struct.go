package main

import "fmt"

// Syntax
// type <name of struct> struct {
// 	//members
// }

type employee struct {
	empid       int
	name        string
	designation string
	department  string
}

func main() {
	fmt.Println("Demo: structure")

	var emp employee = employee{empid: 1, name: "Anand", designation: "Software engineer", department: "MDE"}

	fmt.Println("Empid is :", emp.empid)
	fmt.Println("Emp name:", emp.name)
	fmt.Println("designation:", emp.designation)
	fmt.Println("department:", emp.department)

	emp.designation = emp.department + ":" + emp.designation

	fmt.Println(emp)
}
