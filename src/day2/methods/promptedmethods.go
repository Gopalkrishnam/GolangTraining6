package main

import "fmt"

type Department struct {
	deptId      int
	name        string
	description string
}

func (d Department) Print() {
	fmt.Println("Department details")
	fmt.Println("DepartmentID", d.deptId)
	fmt.Println("DepartmentName", d.name)
	fmt.Println("Dexcription", d.description)
}

type Emp struct {
	empId             int
	name, designation string
	Department
}

func (e Emp) PrintDetails() {
	fmt.Println("Employee details")
	fmt.Println("ID:", e.empId)
	fmt.Println("Name:", e.name)
	fmt.Println("Designation", e.designation)
	fmt.Println(e.Department)
}

func main() {
	fmt.Println("Demo: Prompted methods")

	emp := Emp{empId: 1, name: "Arun kumar", designation: "Engineer", Department: Department{deptId: 1, name: "MDE", description: "Deals with mobile digital enablement"}}

	emp.PrintDetails()

	emp.Print()

}
