package main

import "fmt"

type Student struct {
	RollNo int
	Name   string
	Class  string
}

type Marks struct {
	std                Student
	s1, s2, s3, s4, s5 int
}

func main() {
	fmt.Println("Demo:nested structure")
	var m Marks = Marks{std: Student{RollNo: 1, Name: "Vijay", Class: "10B"}, s1: 100, s2: 99, s3: 98, s4: 100, s5: 99}

	fmt.Println("Student details:", m.std.RollNo, m.std.Name, m.std.Class)
	fmt.Println("Marska", m.s1, m.s2, m.s3, m.s4, m.s5)
	fmt.Println("Total:(", m.s1+m.s2+m.s3+m.s4+m.s5, ")/500")

}
