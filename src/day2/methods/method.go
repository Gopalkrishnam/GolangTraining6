package main

import "fmt"

type Student struct {
	roll_no int
	name    string
	class   string
}

func (std Student) Print() {
	fmt.Println("Details of Student..")
	fmt.Println("RollNo :", std.roll_no)
	fmt.Println("Name :", std.name)
	fmt.Println("Class :", std.class)
}

func (std *Student) Update(rollno int, name, class string) {
	std.name = name
	std.class = class
	std.roll_no = rollno
}

func main() {
	fmt.Println("Methods")

	var student Student = Student{roll_no: 1, name: "Anand", class: "10B"}

	student1 := Student{roll_no: 2, name: "Vijay", class: "10B"}

	student2 := Student{roll_no: 3, name: "Anil", class: "8B"}

	student1.Print()
	student2.Print()
	student.Print()

	var std Student
	std.Update(1, "Vasant", "8B")

	std.Print()

}
