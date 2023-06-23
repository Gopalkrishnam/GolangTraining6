/*
		{
		"name": "Anil kumar",
		"roll_no": 1,
		"Class": "10b"
	}
*/
package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name   string `json:"name"`
	RollNo int    `json:"roll_no"`
	Class  string `json:"Class"`
}

func main() {
	fmt.Println("Json handling...")
	var std = Student{}
	jsonstring := `{
		"name": "Anil kumar",
		"roll_no": 1,
		"Class": "10b"
	}`

	err := json.Unmarshal([]byte(jsonstring), &std)

	if err != nil {
		fmt.Println("Received error while unmarshling", err)
	} else {
		fmt.Println(std)

	}

	data, err := json.Marshal(std)
	if err != nil {
		fmt.Println("Received error while marshling...")
	} else {
		fmt.Println(string(data))
	}

	data, err = json.MarshalIndent(std, "\n", "\t")
	if err != nil {
		fmt.Println("Received error while marshling...")
	} else {
		fmt.Println(string(data))
	}

}
