package fil

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

var id = 0

type TODO struct {
	Task         string    `json:"task"`
	Desc         string    `json:"desc"`
	Done         bool      `json:"done"`
	Created_date time.Time `json:"creat"`
}
type TODOS struct {
    Id int `json:"id"`
	Todos []TODO
}

func Mtodo() {
	what := ""
	fmt.Scan(&what)
	switch {
	case what == "add":
		Add()
	case what == "rm":
		fmt.Println("todo")
	case what == "show":
		Show()
	}
}
func Show() {
	file, err := ioutil.ReadFile("todos.json")
	if err != nil {
		print(err)
	}
	var n TODO
	err = json.Unmarshal(file, &n)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(file))
}
func Add() {

	new_task := ""
	new_desc := ""
	fmt.Println("Name of the new Task: ")
	fmt.Scan(&new_task)
	fmt.Println("Description of the", new_task)
	fmt.Scan(&new_desc)

	file, err := ioutil.ReadFile("todos.json")
	if err != nil {
		print(err)
	}

	data := []TODOS{}

	json.Unmarshal(file, &data)

	new_data := &TODOS{
		Id: 10,
		Todos: []TODO{
			TODO{
				Done:         false,
				Task:         new_task,
				Desc:         new_desc,
				Created_date: time.Now(),
			},
		},
	}
	data = append(data, *new_data)

	dataByte, _ := json.MarshalIndent(data, "", "\t")

	_ = ioutil.WriteFile("todos.json", dataByte, 0644)
}
