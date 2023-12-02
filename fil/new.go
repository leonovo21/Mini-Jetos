package fil

import (
	"fmt"
)

func Nmain(){
	var s string
	fmt.Println("Whats your name ?")
	fmt.Scanln(&s)
	fmt.Println("Hello",s)

}
