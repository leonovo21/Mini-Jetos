package fil

import "fmt"

func Wmain() {
	var s, x = false, 0
	for s == false {
		go worker(s, x)
		go worker(s, x)
		go worker(s, x)
	}
}
func worker(b bool, x int) {
	for b == false {
		fmt.Println(x)
		x++
	}
	return
}
