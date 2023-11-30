package fil

import "fmt"

func Mslacker() {
	age := 0
	fmt.Println("Whats you age? ")
	fmt.Scan(&age)
	cal(age)

}

func cal(age int) (ok bool) {
	if ok {
		fmt.Printf("%d is Old enought\n", age)
	} else {
		fmt.Printf("%d is not Old enought\n", age)
	}

	if age >= 18 {
		return ok
	} else {
		return false
	}
}
