package fil

import (
	"encoding/json"
	"io/ioutil"
)

type Salary struct {
	Basic, HRA, TA float64
}

type Employee struct {
	Task, Desc, Email string
	Age               int
	MonthlySalary     []Salary
}

func Mt() {
	data := Employee{
		Task:  "Mark",
		Desc:  "Jones",
		Email: "mark@gmail.com",
		Age:   25,
		MonthlySalary: []Salary{
			Salary{
				Basic: 15000.00,
				HRA:   5000.00,
				TA:    2000.00,
			},
			Salary{
				Basic: 16000.00,
				HRA:   5000.00,
				TA:    2100.00,
			},
			Salary{
				Basic: 17000.00,
				HRA:   5000.00,
				TA:    2200.00,
			},
		},
	}

	file, _ := json.MarshalIndent(data, "", " ")

	_ = ioutil.WriteFile("test.json", file, 0644)
}
