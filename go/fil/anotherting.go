package fil

import (
	"fmt"
)

var email, pass, choice = "", "", ""
var Inserted_email, Inserted_pass = "", ""
var pass_and_email_exist bool

func ASmain() {
	fmt.Println("Whats do you want to do")
	fmt.Scan(&choice)
	switch {
	case choice == "Create":
		Creat_Acount()
	case choice == "Login":
		Check_Login()
	}
}
func Creat_Acount() {
	fmt.Println("Your Email: ")
	fmt.Scan(&email)
	fmt.Println("Your Password: ")
	fmt.Scan(&pass)
	ASmain()
}
func Check_Email_Exist() {

	if email != "" && pass != "" {
		pass_and_email_exist = true
	}

}
func Check_Login() {
	Check_Email_Exist()
	fmt.Println("Whats your Email?")
	fmt.Scan(&Inserted_email)
	fmt.Println("Whats your Password?")
	fmt.Scan(&Inserted_pass)
	if pass_and_email_exist == true {
		if Inserted_email == email && Inserted_pass == pass {
			fmt.Println("You have logged succefolly")
			ASmain()
		} else {
			fmt.Println("Email or Password incorrect")
			ASmain()
		}
	} else {
		fmt.Println("No email or pass word to log in")
		ASmain()
	}
}
