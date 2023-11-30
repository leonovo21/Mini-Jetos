package math

import (
	"fmt"
)

func media(i int, a int, d int) int {
	fmt.Print("Type a tree numbers: ")
	fmt.Scan(&i, &a, &d)

	fmt.Print("The media is: ", (a+i+d)/3)
	return a
}

func main() {
	media(0, 0, 0)
}
