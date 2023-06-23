package main

import (
	"encoding/xml"
	"fmt"
)

type Employee struct {
	Eid         int    `xml:"id"`
	Name        string `xml:"name"`
	Dept        string `xml:"department"`
	Designation string `xml:"Designation"`
}

func main() {

	var emp = Employee{Eid: 1, Name: "Vijay Kumar", Dept: "Radio", Designation: "Engineer"}

	bytes, err := xml.Marshal(emp)
	if err != nil {
		fmt.Println("Received error while marshling xml")
	} else {
		fmt.Println(string(bytes))
	}
	bytes, err = xml.MarshalIndent(emp, "\n", "    ")
	if err != nil {
		fmt.Println("Received error while marshling xml")
	} else {
		fmt.Println(string(bytes))
	}

	xmlstring := `<Employee>

    <id>1</id>

    <name>Vijay Kumar</name>

    <department>Radio</department>

    <Designation>Engineer</Designation>

</Employee>`
	var e = Employee{}
	err = xml.Unmarshal([]byte(xmlstring), &e)
	if err != nil {
		fmt.Println("Recevied error")
	} else {
		fmt.Println(e)
	}
}
