package fil

import (
	"fmt"
)

func Ymain() {
	var q, n int
	var list []int
	fmt.Println("How many number do you want?: ")
	fmt.Scanf("%d", &q)
	fmt.Println(q)
	for i := 0; i < q; i++{
		fmt.Scanf("%d", &n)
		list = append(list, n)
	}
	fmt.Println(list)

}
