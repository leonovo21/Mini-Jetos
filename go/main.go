package main

import (
	"fmt"
	worker "pro/fil"
)

func main() {
	choices := `				====================================
				ASCII	| ascii | NewHello| nhello |
				Eternal	| worker| HashMap | hash   | 
				hoi3	| hoi3	| Todo	  | todo   |
				New	| new   | Server  | ser    |
                		====================================`
	fmt.Println(choices)
	choice := ""
	fmt.Scan(&choice)

	switch {
	case choice == "new":
		worker.Nmain()
	case choice == "nhello":
		worker.Mhello()
	case choice == "hash":
		worker.Hmain()
	case choice == "worker":
		worker.Wmain()
	case choice == "ascii":
		worker.Macci()
	case choice == "hoi3":
		worker.Hoi3()
	case choice == "todo":
		worker.Mtodo()
	case choice == "ser":
		worker.Smain()
	}
}
