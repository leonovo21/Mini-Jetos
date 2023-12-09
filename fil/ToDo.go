package fil

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

var id = 0

type TODO struct {
	Task          string    `json:"task"`
	Done          bool      `json:"done"`
	Desc          string    `json:"desc"`
	Created_date  time.Time `json:"creat"`
	Finished_date time.Time `json:"finished"`
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
	jsonFile, _ := os.Open("todos.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var todos TODO

	json.Unmarshal(byteValue, &todos)

	fmt.Println(todos.Task)
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

	data := []TODO{}

	json.Unmarshal(file, &data)

	new_data := &TODO{
				Done:         false,
				Task:         new_task,
				Desc:         new_desc,
				Created_date: time.Now(),
			}
	data = append(data, *new_data)

	dataByte, _ := json.MarshalIndent(data, "", "")

	_ = ioutil.WriteFile("todos.json", dataByte, 0644)
}
