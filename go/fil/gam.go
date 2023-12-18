package fil

import (
	"fmt"
)


type Class struct {
    Name string
	Gens []Genes
	Stats []stats

}
type Genes struct {
    Hair string
	Gender string
	Height int
	Frame string
}
type stats struct {
    Strengh int
	Dexterty int
	Intelect int

}
func Gmain(){
	Name := ""
	fmt.Println("Whats your name")
	fmt.Scan(&Name)

	Character_data := Class{
		Name: Name,
		Gens: []Genes{
			Genes{
				Hair: "new",
				Gender: "male",
				Height: 20,
				Frame: "Small",
			},
		},
		Stats: []stats{
			stats{
				Strengh: 20,
				Dexterty: 20,
				Intelect: 20,
			},
		},
	}
	fmt.Println(Character_data)
}
