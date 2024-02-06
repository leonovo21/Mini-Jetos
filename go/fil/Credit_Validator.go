package fil

import (
	"fmt"
)

var CreditCard_Number, WorkingCreditCard, sum, count int
var result string

func M_Validator() {

	for CreditCard_Number <= 0 {
		fmt.Println("What's Your Credit Card Number")
		fmt.Scan(&CreditCard_Number)
	}
	WorkingCreditCard = CreditCard_Number

	for WorkingCreditCard > 0 {
		lastDigit := WorkingCreditCard % 10
		sum = sum + lastDigit
		WorkingCreditCard = WorkingCreditCard / 100
	}

	WorkingCreditCard = CreditCard_Number / 10

	for WorkingCreditCard > 0 {
		lastDigit := WorkingCreditCard % 10
		timesTwo := lastDigit * 2
		sum = sum + (timesTwo % 10) + (timesTwo / 10)
		WorkingCreditCard = WorkingCreditCard / 100
	}
	WorkingCreditCard = CreditCard_Number
	for WorkingCreditCard != 0 {
		WorkingCreditCard = WorkingCreditCard / 10
		count++
	}
	divisor := 10
	for i := 0; i < count-2; i++ {
		divisor = divisor * 10
	}
	firstDigit := CreditCard_Number
	firstTwoDigits := CreditCard_Number / (divisor / 10)

	if sum%10 == 0 {
		if firstDigit == 4 && count == 13 || count == 16 {
			fmt.Println(result, "Visa")
		} else if firstTwoDigits == 34 || firstTwoDigits == 37 && count == 15 {
			fmt.Println(result, "AMEX")
		} else if firstTwoDigits < 50 && firstTwoDigits < 56 && count == 16 {
			fmt.Println(result, "MASTERCARD")
		} else {
			fmt.Println(result, "INVALID")
		}
	} else {
		fmt.Println(result, "INVALID")
	}
	fmt.Println(result)
}
