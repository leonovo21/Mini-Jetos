package fil

import (
	"fmt"
	"time"
)

func Mhello() {
	greenText := " "
	var target string
	ttp := "           "
	index := 0
	fmt.Println("What is your word of choice?")
	fmt.Scan(&target)

	for _, e := range target {
		if e == ' ' {
			index++
			continue
		}
		for _, x := range "abcdefghijklmnopqrstuvwxyz" {
			t := []rune(ttp)
			t[index] = x
			ttp = string(t)
			ttx := fmt.Sprintf("%s%s", greenText, ttp)

			time.Sleep(100 * time.Millisecond)
			fmt.Print("")
			fmt.Println(ttx)
			if e == x {
				index++
				break
			}
		}
	}

}
