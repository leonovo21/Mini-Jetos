package fil

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type TODO struct {
	Task          string    `json:"task"`
	Done          bool      `json:"done"`
	Desc          string    `json:"desc"`
	Created_date  time.Time `json:"creat"`
	Finished_date time.Time `json:"finished"`
}
type TODOS struct {
	Id        int `json:"Id"`
	Task_Name []TODO
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
	content, err := os.ReadFile("todos.json")
	if err != nil {
		fmt.Println("No todos")
	}
	fmt.Println(string(content))
}
func Add() {

	new_task := ""
	new_desc := ""
	fmt.Println("Name of the new Task: ")
	fmt.Scan(&new_task)
	fmt.Println("Description of the", new_task)
	fmt.Scan(&new_desc)
	T := TODOS{
		Id: +1,
		Task_Name: []TODO{
			TODO{
				Done:         false,
				Task:         new_task,
				Desc:         new_desc,
				Created_date: time.Now(),
			},
		},
	}

	file, _ := json.MarshalIndent(T, "", "")

	_ = ioutil.WriteFile("todos.json", file, 0644)
}
