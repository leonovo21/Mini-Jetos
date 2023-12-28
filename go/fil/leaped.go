package fil

import (
	"fmt"
)

// var d = bool
func Lmain() {
	y := 0
	fmt.Printf("Whats year do you wish to check: ")
	fmt.Scanf("%d", &y)
	Leaped(y)
	// fmt.Println(d)
}
func Leaped(y int) bool {
	if y%4 == 0 {
		if y%100 == 0 {
			if y%400 == 0 {
				print("true")
				return true
			} else {
				print("false")
				return false
			}
		}
		print("true")
		return true
	} else {
		print("false")
		return false
	}

}
