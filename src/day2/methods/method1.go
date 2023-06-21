package main

import "fmt"

type Emp struct {
	empId                         int
	name, department, designation string
}

func (e Emp) Print() {
	fmt.Println("Employee details")
	fmt.Println("EmpID:", e.empId)
	fmt.Println("Name:", e.name)
	fmt.Println("Department:", e.department)
	fmt.Println("Designation:", e.designation)
}

type Department struct {
	emplist     []Emp
	deptId      int
	name        string
	description string
}

func (d Department) Print() {
	fmt.Println("Department details")
	fmt.Println("DepartmentID", d.deptId)
	fmt.Println("DepartmentName", d.name)
	fmt.Println("Dexcription", d.description)
	for _, val := range d.emplist {
		val.Print()
	}
}

func main() {
	fmt.Println("Methods")

	var empList = []Emp{Emp{empId: 1, name: "Anand", department: "Radio", designation: "Engineer"},
		Emp{empId: 2, name: "Arun", department: "Radio", designation: "Engineer"}}

	var dept = Department{deptId: 1, name: "Radio", description: "Works on mobile radio", emplist: empList}

	dept.Print()
}
